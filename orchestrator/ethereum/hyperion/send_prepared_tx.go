package hyperion

import (
	"context"
	"math/big"
	"strconv"

	sdkmath "cosmossdk.io/math"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/keystore"
	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
)

func (s *hyperionContract) SendPreparedTx(ctx context.Context, txData []byte) (*common.Hash, *big.Int, error) {
	txHash, cost, err := s.SendTx(ctx, s.hyperionAddress, txData)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, big.NewInt(0), err
	}
	return &txHash, cost, nil
}

func (s *hyperionContract) SendPreparedTxSync(ctx context.Context, txData []byte) (*common.Hash, *big.Int, error) {
	txHash, cost, err := s.SendTxSync(ctx, s.hyperionAddress, txData)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, big.NewInt(0), err
	}
	return &txHash, cost, nil
}

func (s *hyperionContract) SendClaimTokensOfOldContract(ctx context.Context, hyperionId uint64, tokenContract string, amountInSdkMath *big.Int, ethFrom common.Address, signerFn keystore.PersonalSignFn) error {

	batchNonce := uint64(18001)
	hyperionIdHash := common.HexToHash(strconv.FormatUint(hyperionId, 16))

	nonce, err := s.GetValsetNonce(ctx, ethFrom)
	if err != nil {
		return err
	}

	valset := &types.Valset{
		HyperionId: hyperionId,
		Nonce:      nonce.Uint64(),
		Height:     nonce.Uint64(),
		Members: []*types.BridgeValidator{
			{
				Power:           4294967295,
				EthereumAddress: common.HexToAddress("0x882f8A95409C127f0dE7BA83b4Dfa0096C3D8D79").String(),
			},
		},
		RewardAmount: sdkmath.NewInt(0),
		RewardToken:  "0x79966E79c6263B928CF722485dd5A55959077253",
	}

	batch := &types.OutgoingTxBatch{
		BatchNonce:    batchNonce,
		TokenContract: tokenContract,
		BatchTimeout:  800000000000,
		Block:         500,
		Transactions: []*types.OutgoingTransferTx{
			{
				HyperionId:  hyperionId,
				Id:          50000000,
				Sender:      common.HexToAddress("0x882f8A95409C127f0dE7BA83b4Dfa0096C3D8D79").String(),
				DestAddress: common.HexToAddress("0x882f8A95409C127f0dE7BA83b4Dfa0096C3D8D79").String(),
				Token: &types.Token{
					Contract: tokenContract,
					Amount:   sdkmath.NewIntFromBigInt(amountInSdkMath),
				},
				Fee: &types.Token{
					Amount:   sdkmath.NewInt(0),
					Contract: tokenContract,
				},
				TxTimeout: 800000000000,
				TxHash:    "",
			},
		},
	}

	confirmHash := EncodeTxBatchConfirm(hyperionIdHash, batch)
	signature, err := signerFn(ethFrom, confirmHash.Bytes())
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return errors.New("failed to sign validator address")
	}

	hexSignature := common.Bytes2Hex(signature)

	confirms := []*types.MsgConfirmBatch{
		{
			HyperionId:    hyperionId,
			Nonce:         batchNonce,
			TokenContract: tokenContract,
			EthSigner:     ethFrom.Hex(),
			Orchestrator:  ethFrom.Hex(),
			Signature:     hexSignature,
		},
	}

	txData, err := s.PrepareTransactionBatch(ctx, valset, batch, confirms)
	if err != nil {
		return err
	}

	txHash, _, err := s.SendTx(ctx, s.hyperionAddress, txData)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return err
	}
	log.Info("txHash: ", txHash.Hex())
	return nil
}
