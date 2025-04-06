package hyperion

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/committer"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	wrappers "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
)

type HyperionContract interface {
	committer.EVMCommitter

	Address() common.Address

	SendToHelios(
		ctx context.Context,
		erc20 common.Address,
		amount *big.Int,
		cosmosAccAddress sdk.AccAddress,
		senderAddress common.Address,
		data string,
	) (*common.Hash, error)

	SendTransactionBatch(
		ctx context.Context,
		currentValset *types.Valset,
		batch *types.OutgoingTxBatch,
		confirms []*types.MsgConfirmBatch,
	) (*common.Hash, error)

	SendEthValsetUpdate(
		ctx context.Context,
		oldValset *types.Valset,
		newValset *types.Valset,
		confirms []*types.MsgValsetConfirm,
	) (*common.Hash, error)

	GetTxBatchNonce(
		ctx context.Context,
		erc20ContractAddress common.Address,
		callerAddress common.Address,
	) (*big.Int, error)

	GetValsetNonce(
		ctx context.Context,
		callerAddress common.Address,
	) (*big.Int, error)

	GetHyperionID(
		ctx context.Context,
		callerAddress common.Address,
	) (common.Hash, error)

	GetERC20Symbol(
		ctx context.Context,
		erc20ContractAddress common.Address,
		callerAddress common.Address,
	) (symbol string, err error)

	SubscribeToPendingTxs(
		alchemyWebsocketURL string)

	GetLastEventNonce(
		ctx context.Context,
		callerAddress common.Address,
	) (*big.Int, error)

	GetLastValsetCheckpoint(
		ctx context.Context,
		callerAddress common.Address,
	) (*common.Hash, error)

	GetLastValsetUpdatedEventHeight(
		ctx context.Context,
		callerAddress common.Address,
	) (*big.Int, error)

	GetLastEventHeight(
		ctx context.Context,
		callerAddress common.Address,
	) (*big.Int, error)
}

func NewHyperionContract(
	ethCommitter committer.EVMCommitter,
	hyperionAddress common.Address,
	pendingTxInputList PendingTxInputList,
	pendingTxWaitDuration time.Duration,
) (HyperionContract, error) {
	fmt.Println("Contract hyperionAddress", hyperionAddress.String())
	ethHyperion, err := wrappers.NewHyperion(hyperionAddress, ethCommitter.Provider())
	if err != nil {
		return nil, err
	}

	svc := &hyperionContract{
		EVMCommitter:          ethCommitter,
		hyperionAddress:       hyperionAddress,
		ethHyperion:           ethHyperion,
		pendingTxInputList:    pendingTxInputList,
		pendingTxWaitDuration: pendingTxWaitDuration,
		svcTags: metrics.Tags{
			"svc": "hyperion_contract",
		},
	}

	return svc, nil
}

type hyperionContract struct {
	committer.EVMCommitter

	ethProvider     provider.EVMProvider
	hyperionAddress common.Address
	ethHyperion     *wrappers.Hyperion

	pendingTxInputList    PendingTxInputList
	pendingTxWaitDuration time.Duration

	svcTags metrics.Tags
}

func (s *hyperionContract) Address() common.Address {
	return s.hyperionAddress
}

// maxUintAllowance is uint constant MAX_UINT = 2**256 - 1
var maxUintAllowance = big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))

var (
	hyperionABI, _ = abi.JSON(strings.NewReader(wrappers.HyperionABI))
	erc20ABI, _    = abi.JSON(strings.NewReader(wrappers.ERC20ABI))
)

func sigToVRS(sigHex string) (v uint8, r, s common.Hash) {
	signatureBytes := common.FromHex(sigHex)
	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}

	v = vParam
	r = common.BytesToHash(signatureBytes[0:32])
	s = common.BytesToHash(signatureBytes[32:64])

	return
}

// The total power in the Hyperion bridge is normalized to u32 max every
// time a validator set is created. This value of up to u32 max is then
// stored in a i64 to prevent overflow during computation.
const totalHyperionPower int64 = math.MaxUint32

// hyperionPowerToPercent takes in an amount of power in the hyperion bridge, returns a percentage of total
func hyperionPowerToPercent(total *big.Int) float32 {
	d := decimal.NewFromBigInt(total, 0)
	f, _ := d.Div(decimal.NewFromInt(totalHyperionPower)).Shift(2).Float64()
	return float32(f)
}

var ErrInsufficientVotingPowerToPass = errors.New("insufficient voting power")
