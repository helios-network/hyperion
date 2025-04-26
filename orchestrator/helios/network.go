package helios

import (
	"context"
	"fmt"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/keystore"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios/hyperion"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios/tendermint"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/Helios-Chain-Labs/sdk-go/client/chain"
	clientcommon "github.com/Helios-Chain-Labs/sdk-go/client/common"
	comethttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type NetworkConfig struct {
	ChainID,
	ValidatorAddress,
	HeliosGRPC,
	TendermintRPC,
	GasPrice string
	Gas string
}

type NetworkWithoutBroadcast interface {
	hyperion.QueryClient
	tendermint.Client
}

type Network interface {
	hyperion.QueryClient
	hyperion.BroadcastClient
	tendermint.Client
}

func NewNetworkWithoutBroadcast(k keyring.Keyring, cfg NetworkConfig) (NetworkWithoutBroadcast, error) {
	clientCfg := cfg.loadClientConfig()

	log.Infoln("New Client Context with chain", clientCfg.ChainId, " and Validator", cfg.ValidatorAddress)

	clientCtx, err := chain.NewClientContext(clientCfg.ChainId, cfg.ValidatorAddress, k)
	if err != nil {
		return nil, err
	}

	log.Infoln("Context OK")

	log.Infoln("NodeURI", clientCfg.TmEndpoint)

	clientCtx.WithNodeURI(clientCfg.TmEndpoint)

	log.Infoln("Node URI OK")

	tmRPC, err := comethttp.New(clientCfg.TmEndpoint, "/websocket")
	if err != nil {
		return nil, err
	}

	log.Infoln("RPC OK")

	clientCtx = clientCtx.WithClient(tmRPC)

	log.Infoln("WithClient OK")

	log.Infoln(fmt.Sprintf("GasPrice CONFIG %s", cfg.GasPrice))
	log.Infoln(fmt.Sprintf("GAS CONFIG %s", cfg.Gas))

	chainClient, err := chain.NewChainClient(clientCtx, clientCfg, clientcommon.OptionGasPrices(cfg.GasPrice), clientcommon.OptionGas(cfg.Gas))
	if err != nil {
		return nil, err
	}

	log.Infoln("NewChainClient OK")

	time.Sleep(1 * time.Second)

	conn := awaitConnection(chainClient, 1*time.Minute)

	log.Infoln("CON OK")

	net := struct {
		hyperion.QueryClient
		tendermint.Client
	}{
		hyperion.NewQueryClient(hyperiontypes.NewQueryClient(conn)),
		tendermint.NewRPCClient(clientCfg.TmEndpoint),
	}

	log.Infoln("NET LOADED")

	return net, nil
}

func NewNetworkWithBroadcast(k keyring.Keyring, ethSignFn keystore.PersonalSignFn, cfg NetworkConfig) (Network, error) {
	clientCfg := cfg.loadClientConfig()

	log.Infoln("New Client Context with chain", clientCfg.ChainId, " and Validator", cfg.ValidatorAddress)

	clientCtx, err := chain.NewClientContext(clientCfg.ChainId, cfg.ValidatorAddress, k)
	if err != nil {
		return nil, err
	}

	log.Infoln("Context OK")

	log.Infoln("NodeURI", clientCfg.TmEndpoint)

	clientCtx.WithNodeURI(clientCfg.TmEndpoint)

	log.Infoln("Node URI OK")

	tmRPC, err := comethttp.New(clientCfg.TmEndpoint, "/websocket")
	if err != nil {
		return nil, err
	}

	log.Infoln("RPC OK")

	clientCtx = clientCtx.WithClient(tmRPC)

	log.Infoln("WithClient OK")

	log.Infoln(fmt.Sprintf("GasPrice CONFIG %s", cfg.GasPrice))
	log.Infoln(fmt.Sprintf("GAS CONFIG %s", cfg.Gas))

	chainClient, err := chain.NewChainClient(clientCtx, clientCfg, clientcommon.OptionGasPrices(cfg.GasPrice), clientcommon.OptionGas(cfg.Gas))
	if err != nil {
		return nil, err
	}

	log.Infoln("NewChainClient OK")

	time.Sleep(1 * time.Second)

	conn := awaitConnection(chainClient, 1*time.Minute)

	log.Infoln("CON OK")

	net := struct {
		hyperion.QueryClient
		hyperion.BroadcastClient
		tendermint.Client
	}{
		hyperion.NewQueryClient(hyperiontypes.NewQueryClient(conn)),
		hyperion.NewBroadcastClient(chainClient, ethSignFn),
		tendermint.NewRPCClient(clientCfg.TmEndpoint),
	}

	log.Infoln("NET LOADED")

	return net, nil
}

func awaitConnection(client chain.ChainClient, timeout time.Duration) *grpc.ClientConn {
	ctx, cancelWait := context.WithTimeout(context.Background(), timeout)
	defer cancelWait()

	grpcConn := client.QueryClient()

	for {
		select {
		case <-ctx.Done():
			log.Fatalln("GRPC service wait timed out")
		default:
			state := grpcConn.GetState()
			if state != connectivity.Ready {
				log.WithField("state", state.String()).Warningln("state of GRPC connection not ready")
				time.Sleep(5 * time.Second)
				continue
			}

			return grpcConn
		}
	}
}

func (cfg NetworkConfig) loadClientConfig() clientcommon.Network {
	if custom := cfg.HeliosGRPC != "" && cfg.TendermintRPC != ""; custom {
		log.WithFields(log.Fields{"helios_grpc": cfg.HeliosGRPC, "tendermint_rpc": cfg.TendermintRPC}).Debugln("using custom endpoints for Helios")
		return customEndpoints(cfg)
	}

	c := loadBalancedEndpoints(cfg)
	log.WithFields(log.Fields{"Helios_grpc": c.ChainGrpcEndpoint, "tendermint_rpc": c.TmEndpoint}).Debugln("using load balanced endpoints for Helios")

	return c
}

func customEndpoints(cfg NetworkConfig) clientcommon.Network {
	c := clientcommon.LoadNetwork("devnet", "")
	c.Name = "custom"
	c.ChainId = cfg.ChainID
	c.FeeDenom = "helios"
	c.TmEndpoint = cfg.TendermintRPC
	c.ChainGrpcEndpoint = cfg.HeliosGRPC
	c.ExplorerGrpcEndpoint = ""
	c.LcdEndpoint = ""
	c.ExplorerGrpcEndpoint = ""

	return c
}

func loadBalancedEndpoints(cfg NetworkConfig) clientcommon.Network {
	var networkName string
	switch cfg.ChainID {
	case "42000":
		networkName = "mainnet"
	case "42001":
		networkName = "testnet"
	case "42002":
		networkName = "devnet"
	default:
		panic(fmt.Errorf("no provider for chain id %s", cfg.ChainID))
	}

	return clientcommon.LoadNetwork(networkName, "lb")
}

func HasRegisteredOrchestrator(n Network, hyperionId uint64, ethAddr gethcommon.Address) (cosmostypes.AccAddress, bool) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	log.Info("ethAddr: ", ethAddr)

	validator, err := n.GetValidatorAddress(ctx, hyperionId, ethAddr)
	log.Info("validator: ", validator)
	if err != nil {
		return nil, false
	}

	return validator, true
}
