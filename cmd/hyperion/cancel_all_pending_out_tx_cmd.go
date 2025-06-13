package main

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/cmd/hyperion/queries"
	globaltypes "github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/gorilla/mux"
	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
)

func cancelAllPendingOutTxCmd(cmd *cli.Cmd) {
	cmd.Before = func() {
		initMetrics(cmd)
	}

	cfg := initConfig(cmd)

	cmd.Action = func() {
		// ensure a clean exit
		defer closer.Close()

		router := mux.NewRouter()
		router.Use(loggingMiddleware)

		global := globaltypes.NewGlobal(&globaltypes.Config{
			PrivateKey:            *cfg.heliosPrivKey,
			HeliosChainID:         *cfg.heliosChainID,
			HeliosGRPC:            *cfg.heliosGRPC,
			TendermintRPC:         *cfg.tendermintRPC,
			HeliosGasPrices:       *cfg.heliosGasPrices,
			HeliosGas:             *cfg.heliosGas,
			EthGasPriceAdjustment: *cfg.ethGasPriceAdjustment,
			EthMaxGasPrice:        *cfg.ethMaxGasPrice,
			PendingTxWaitDuration: *cfg.pendingTxWaitDuration,
		})
		global.InitHeliosNetwork(0)

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		queries.CancelAllPendingOutTx(ctx, global, 80002)

		closer.Hold()
	}
}
