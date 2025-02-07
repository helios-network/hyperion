package main

import (
	"context"
	"time"

	"github.com/Helios-Chain-Labs/peggo/orchestrator/cosmos/hyperion"

	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/peggo/orchestrator/cosmos"
)

// txCmdSubset contains actions that can sign and send messages to Cosmos module
// as well as Ethereum transactions to Hyperion contract.
//
// $ peggo tx
func txCmdSubset(cmd *cli.Cmd) {
	cmd.Command(
		"register-eth-key",
		"Submits an Ethereum key that will be used to sign messages on behalf of your Validator",
		registerEthKeyCmd,
	)
}

func registerEthKeyCmd(cmd *cli.Cmd) {
	var (
		// Cosmos params
		cosmosChainID   *string
		cosmosGRPC      *string
		tendermintRPC   *string
		cosmosGasPrices *string
		cosmosGas       *string

		// Cosmos Key Management
		cosmosKeyringDir     *string
		cosmosKeyringAppName *string
		cosmosKeyringBackend *string

		cosmosKeyFrom       *string
		cosmosKeyPassphrase *string
		cosmosPrivKey       *string
		cosmosUseLedger     *bool

		// Ethereum Key Management
		ethKeystoreDir *string
		ethKeyFrom     *string
		ethPassphrase  *string
		ethPrivKey     *string
		ethUseLedger   *bool

		// Misc
		alwaysAutoConfirm *bool
	)

	initCosmosOptions(
		cmd,
		&cosmosChainID,
		&cosmosGRPC,
		&tendermintRPC,
		&cosmosGasPrices,
		&cosmosGas,
	)

	initCosmosKeyOptions(
		cmd,
		&cosmosKeyringDir,
		&cosmosKeyringAppName,
		&cosmosKeyringBackend,
		&cosmosKeyFrom,
		&cosmosKeyPassphrase,
		&cosmosPrivKey,
		&cosmosUseLedger,
	)

	initEthereumKeyOptions(
		cmd,
		&ethKeystoreDir,
		&ethKeyFrom,
		&ethPassphrase,
		&ethPrivKey,
		&ethUseLedger,
	)

	initInteractiveOptions(
		cmd,
		&alwaysAutoConfirm,
	)

	cmd.Action = func() {
		// ensure a clean exit
		defer closer.Close()

		if *ethUseLedger {
			log.Warningln("beware: you cannot really use Ledger for orchestrator, so make sure the Ethereum key is accessible outside of it")
		}

		keyringCfg := cosmos.KeyringConfig{
			KeyringDir:     *cosmosKeyringDir,
			KeyringAppName: *cosmosKeyringAppName,
			KeyringBackend: *cosmosKeyringBackend,
			KeyFrom:        *cosmosKeyFrom,
			KeyPassphrase:  *cosmosKeyPassphrase,
			PrivateKey:     *cosmosPrivKey,
			UseLedger:      *cosmosUseLedger,
		}

		keyring, err := cosmos.NewKeyring(keyringCfg)
		orShutdown(err)

		ethKeyFromAddress, _, personalSignFn, err := initEthereumAccountsManager(
			0,
			ethKeystoreDir,
			ethKeyFrom,
			ethPassphrase,
			ethPrivKey,
			ethUseLedger,
		)
		if err != nil {
			log.WithError(err).Fatalln("failed to init Ethereum account")
		}

		log.Infoln("Using Cosmos ValAddress", keyring.Addr.String())
		log.Infoln("Using Ethereum address", ethKeyFromAddress.String())

		actionConfirmed := *alwaysAutoConfirm || stdinConfirm("Confirm UpdateHyperionOrchestratorAddresses transaction? [y/N]: ")
		if !actionConfirmed {
			return
		}

		net, err := cosmos.NewNetwork(keyring, personalSignFn, cosmos.NetworkConfig{
			ChainID:          *cosmosChainID,
			ValidatorAddress: keyring.Addr.String(),
			CosmosGRPC:       *cosmosGRPC,
			TendermintRPC:    *tendermintRPC,
			GasPrice:         *cosmosGasPrices,
			Gas:              *cosmosGas,
		})

		if err != nil {
			log.Fatalln("failed to connect to Helios network")
		}

		broadcastCtx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFn()

		if err = hyperion.BroadcastClient(net).UpdateHyperionOrchestratorAddresses(broadcastCtx, ethKeyFromAddress, keyring.Addr); err != nil {
			log.WithError(err).Errorln("failed to broadcast Tx")
			time.Sleep(time.Second)
			return
		}

		log.Infof("Registered Ethereum address %s for validator address %s",
			ethKeyFromAddress, keyring.Addr.String())
	}
}
