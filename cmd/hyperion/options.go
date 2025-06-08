package main

import (
	cli "github.com/jawher/mow.cli"
)

// initGlobalOptions defines some global CLI options, that are useful for most parts of the app.
// Before adding option to there, consider moving it into the actual Cmd.
func initGlobalOptions(
	envName **string,
	appLogLevel **string,
	svcWaitTimeout **string,
) {
	*envName = app.String(cli.StringOpt{
		Name:   "e env",
		Desc:   "The environment name this app runs in. Used for metrics and error reporting.",
		EnvVar: "HYPERION_ENV",
		Value:  "local",
	})

	*appLogLevel = app.String(cli.StringOpt{
		Name:   "l log-level",
		Desc:   "Available levels: error, warn, info, debug.",
		EnvVar: "HYPERION_LOG_LEVEL",
		Value:  "info",
	})

	*svcWaitTimeout = app.String(cli.StringOpt{
		Name:   "svc-wait-timeout",
		Desc:   "Standard wait timeout for external services (e.g. Cosmos daemon GRPC connection)",
		EnvVar: "HYPERION_SERVICE_WAIT_TIMEOUT",
		Value:  "1m",
	})
}

// initStatsdOptions sets options for StatsD metrics.
func initStatsdOptions(
	cmd *cli.Cmd,
	statsdAgent **string,
	statsdPrefix **string,
	statsdAddr **string,
	statsdStuckDur **string,
	statsdMocking **string,
	statsdDisabled **string,
) {
	*statsdAgent = cmd.String(cli.StringOpt{
		Name:   "statsd-agent",
		Desc:   "Specify StatsD agent.",
		EnvVar: "HYPERION_STATSD_AGENT",
		Value:  "telegraf",
	})

	*statsdPrefix = cmd.String(cli.StringOpt{
		Name:   "statsd-prefix",
		Desc:   "Specify StatsD compatible metrics prefix.",
		EnvVar: "HYPERION_STATSD_PREFIX",
		Value:  "hyperion",
	})

	*statsdAddr = cmd.String(cli.StringOpt{
		Name:   "statsd-addr",
		Desc:   "UDP address of a StatsD compatible metrics aggregator.",
		EnvVar: "HYPERION_STATSD_ADDR",
		Value:  "localhost:8125",
	})

	*statsdStuckDur = cmd.String(cli.StringOpt{
		Name:   "statsd-stuck-func",
		Desc:   "Sets a duration to consider a function to be stuck (e.g. in deadlock).",
		EnvVar: "HYPERION_STATSD_STUCK_DUR",
		Value:  "5m",
	})

	*statsdMocking = cmd.String(cli.StringOpt{
		Name:   "statsd-mocking",
		Desc:   "If enabled replaces statsd client with a mock one that simply logs values.",
		EnvVar: "HYPERION_STATSD_MOCKING",
		Value:  "false",
	})

	*statsdDisabled = cmd.String(cli.StringOpt{
		Name:   "statsd-disabled",
		Desc:   "Force disabling statsd reporting completely.",
		EnvVar: "HYPERION_STATSD_DISABLED",
		Value:  "true",
	})
}

type Config struct {
	// Cosmos params
	heliosChainID   *string
	heliosGRPC      *string
	tendermintRPC   *string
	heliosGasPrices *string
	heliosGas       *string

	// Cosmos Key Management
	heliosPrivKey *string

	// Ethereum params
	ethGasPriceAdjustment *float64
	ethMaxGasPrice        *string

	// Relayer config
	pendingTxWaitDuration *string

	// Batch requester config
	minBatchFeeUSD *float64
}

func initConfig(cmd *cli.Cmd) Config {
	cfg := Config{}

	/** Helios **/

	cfg.heliosChainID = cmd.String(cli.StringOpt{
		Name:   "helios-chain-id",
		Desc:   "Specify Chain ID of the Helios network.",
		EnvVar: "HYPERION_HELIOS_CHAIN_ID",
		Value:  "42000",
	})

	cfg.heliosGRPC = cmd.String(cli.StringOpt{
		Name:   "helios-grpc",
		Desc:   "Helios GRPC querying endpoint",
		EnvVar: "HYPERION_HELIOS_GRPC",
		Value:  "tcp://localhost:9090",
	})

	cfg.tendermintRPC = cmd.String(cli.StringOpt{
		Name:   "tendermint-rpc",
		Desc:   "Tendermint RPC endpoint",
		EnvVar: "HYPERION_TENDERMINT_RPC",
		Value:  "http://localhost:26657",
	})

	cfg.heliosGasPrices = cmd.String(cli.StringOpt{
		Name:   "helios-gas-prices",
		Desc:   "Specify Helios chain transaction fees as DecCoins gas prices",
		EnvVar: "HYPERION_HELIOS_GAS_PRICES",
		Value:  "500000000ahelios", // example: 500000000ahelios
	})

	cfg.heliosGas = cmd.String(cli.StringOpt{
		Name:   "helios-gas",
		Desc:   "Specify Helios chain transaction gas",
		EnvVar: "HYPERION_HELIOS_GAS",
		Value:  "2000000", // example: 2000000
	})

	cfg.heliosPrivKey = cmd.String(cli.StringOpt{
		Name:   "helios-pk",
		Desc:   "Provide a raw Helios account private key of the validator in hex.",
		EnvVar: "HYPERION_HELIOS_PK",
	})

	/** Ethereum **/

	cfg.ethGasPriceAdjustment = cmd.Float64(cli.Float64Opt{
		Name:   "eth-gas-price-adjustment",
		Desc:   "gas price adjustment for Ethereum transactions",
		EnvVar: "HYPERION_ETH_GAS_PRICE_ADJUSTMENT",
		Value:  float64(1.3),
	})

	cfg.ethMaxGasPrice = cmd.String(cli.StringOpt{
		Name:   "eth-max-gas-price",
		Desc:   "Specify Max gas price for Ethereum Transactions in GWei",
		EnvVar: "HYPERION_ETH_MAX_GAS_PRICE",
		Value:  "500gwei",
	})

	/** Batch Requester **/

	cfg.minBatchFeeUSD = cmd.Float64(cli.Float64Opt{
		Name:   "min-batch-fee-usd",
		Desc:   "If set, batch request will create batches only if fee threshold exceeds",
		EnvVar: "HYPERION_MIN_BATCH_FEE_USD",
		Value:  float64(23.3),
	})

	return cfg
}
