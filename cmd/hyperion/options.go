package main

import (
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
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
	heliosKeyringDir     *string
	heliosKeyringAppName *string
	heliosKeyringBackend *string

	heliosKeyFrom       *string
	heliosKeyPassphrase *string
	heliosPrivKey       *string

	// Ethereum params
	hyperionID            *int
	ethChainID            *int
	ethNodeRPCs            *string
	ethNodeAlchemyWS      *string
	ethGasPriceAdjustment *float64
	ethMaxGasPrice        *string

	// Ethereum Key Management
	ethKeystoreDir *string
	ethKeyFrom     *string
	ethPassphrase  *string
	ethPrivKey     *string

	// Relayer config
	relayValsets          *bool
	relayValsetOffsetDur  *string
	relayBatches          *bool
	relayBatchOffsetDur   *string
	pendingTxWaitDuration *string

	// Batch requester config
	minBatchFeeUSD *float64

	coingeckoApi *string

	chainParams *hyperiontypes.CounterpartyChainParams
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
	})

	cfg.heliosGasPrices = cmd.String(cli.StringOpt{
		Name:   "helios-gas-prices",
		Desc:   "Specify Helios chain transaction fees as DecCoins gas prices",
		EnvVar: "HYPERION_HELIOS_GAS_PRICES",
		Value:  "", // example: 500000000ahelios
	})

	cfg.heliosGas = cmd.String(cli.StringOpt{
		Name:   "helios-gas",
		Desc:   "Specify Helios chain transaction gas",
		EnvVar: "HYPERION_HELIOS_GAS",
		Value:  "", // example: 2000000
	})

	cfg.heliosKeyringBackend = cmd.String(cli.StringOpt{
		Name:   "helios-keyring",
		Desc:   "Specify Helios keyring backend (os|file|pass|test|local)",
		EnvVar: "HYPERION_HELIOS_KEYRING",
		Value:  "local",
	})

	cfg.heliosKeyringDir = cmd.String(cli.StringOpt{
		Name:   "helios-keyring-dir",
		Desc:   "Specify Helios keyring dir, if using file keyring.",
		EnvVar: "HYPERION_HELIOS_KEYRING_DIR",
		Value:  "",
	})

	cfg.heliosKeyringAppName = cmd.String(cli.StringOpt{
		Name:   "helios-keyring-app",
		Desc:   "Specify Helios keyring app name.",
		EnvVar: "HYPERION_HELIOS_KEYRING_APP",
		Value:  "hyperion",
	})

	cfg.heliosKeyFrom = cmd.String(cli.StringOpt{
		Name:   "helios-from",
		Desc:   "Specify the Helios validator key name or address. If specified, must exist in keyring, ledger or match the privkey.",
		EnvVar: "HYPERION_HELIOS_FROM",
	})

	cfg.heliosKeyPassphrase = cmd.String(cli.StringOpt{
		Name:   "helios-from-passphrase",
		Desc:   "Specify keyring passphrase, otherwise Stdin will be used.",
		EnvVar: "HYPERION_HELIOS_FROM_PASSPHRASE",
		Value:  "hyperion",
	})

	cfg.heliosPrivKey = cmd.String(cli.StringOpt{
		Name:   "helios-pk",
		Desc:   "Provide a raw Helios account private key of the validator in hex.",
		EnvVar: "HYPERION_HELIOS_PK",
	})

	/** Ethereum **/

	cfg.hyperionID = cmd.Int(cli.IntOpt{
		Name:   "hyperion-id",
		Desc:   "Specify Hyperion ID of the Ethereum network.",
		EnvVar: "HYPERION_ID",
		Value:  0,
	})

	cfg.ethChainID = cmd.Int(cli.IntOpt{
		Name:   "eth-chain-id",
		Desc:   "Specify Chain ID of the Ethereum network.",
		EnvVar: "HYPERION_ETH_CHAIN_ID",
		Value:  42,
	})

	cfg.ethNodeRPCs = cmd.String(cli.StringOpt{
		Name:   "eth-node-http",
		Desc:   "Specify HTTP endpoint for an Ethereum node.",
		EnvVar: "HYPERION_ETH_RPCS",
		Value:  "http://localhost:1317",
	})

	cfg.ethNodeAlchemyWS = cmd.String(cli.StringOpt{
		Name:   "eth-node-alchemy-ws",
		Desc:   "Specify websocket url for an Alchemy ethereum node.",
		EnvVar: "HYPERION_ETH_ALCHEMY_WS",
		Value:  "",
	})

	cfg.ethGasPriceAdjustment = cmd.Float64(cli.Float64Opt{
		Name:   "eth_gas_price_adjustment",
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

	cfg.ethKeystoreDir = cmd.String(cli.StringOpt{
		Name:   "eth-keystore-dir",
		Desc:   "Specify Ethereum keystore dir (Geth-format) prefix.",
		EnvVar: "HYPERION_ETH_KEYSTORE_DIR",
	})

	cfg.ethKeyFrom = cmd.String(cli.StringOpt{
		Name:   "eth-from",
		Desc:   "Specify the from address. If specified, must exist in keystore, ledger or match the privkey.",
		EnvVar: "HYPERION_ETH_FROM",
	})

	cfg.ethPassphrase = cmd.String(cli.StringOpt{
		Name:   "eth-passphrase",
		Desc:   "Passphrase to unlock the private key from armor, if empty then stdin is used.",
		EnvVar: "HYPERION_ETH_PASSPHRASE",
	})

	cfg.ethPrivKey = cmd.String(cli.StringOpt{
		Name:   "eth-pk",
		Desc:   "Provide a raw Ethereum private key of the validator in hex. USE FOR TESTING ONLY!",
		EnvVar: "HYPERION_ETH_PK",
	})

	/** Relayer **/

	cfg.relayValsets = cmd.Bool(cli.BoolOpt{
		Name:   "relay_valsets",
		Desc:   "If enabled, relayer will relay valsets to ethereum",
		EnvVar: "HYPERION_RELAY_VALSETS",
		Value:  false,
	})

	cfg.relayValsetOffsetDur = cmd.String(cli.StringOpt{
		Name:   "relay_valset_offset_dur",
		Desc:   "If set, relayer will broadcast valsetUpdate only after relayValsetOffsetDur has passed from time of valsetUpdate creation",
		EnvVar: "HYPERION_RELAY_VALSET_OFFSET_DUR",
		Value:  "5m",
	})

	cfg.relayBatches = cmd.Bool(cli.BoolOpt{
		Name:   "relay_batches",
		Desc:   "If enabled, relayer will relay batches to ethereum",
		EnvVar: "HYPERION_RELAY_BATCHES",
		Value:  false,
	})

	cfg.relayBatchOffsetDur = cmd.String(cli.StringOpt{
		Name:   "relay_batch_offset_dur",
		Desc:   "If set, relayer will broadcast batches only after relayBatchOffsetDur has passed from time of batch creation",
		EnvVar: "HYPERION_RELAY_BATCH_OFFSET_DUR",
		Value:  "5m",
	})

	cfg.pendingTxWaitDuration = cmd.String(cli.StringOpt{
		Name:   "relay_pending_tx_wait_duration",
		Desc:   "If set, relayer will broadcast pending batches/valsetupdate only after pendingTxWaitDuration has passed",
		EnvVar: "HYPERION_RELAY_PENDING_TX_WAIT_DURATION",
		Value:  "20m",
	})

	/** Batch Requester **/

	cfg.minBatchFeeUSD = cmd.Float64(cli.Float64Opt{
		Name:   "min_batch_fee_usd",
		Desc:   "If set, batch request will create batches only if fee threshold exceeds",
		EnvVar: "HYPERION_MIN_BATCH_FEE_USD",
		Value:  float64(23.3),
	})

	/** Coingecko **/

	cfg.coingeckoApi = cmd.String(cli.StringOpt{
		Name:   "coingecko_api",
		Desc:   "Specify HTTP endpoint for pricefeed api.",
		EnvVar: "HYPERION_COINGECKO_API",
		Value:  "https://api.coingecko.com/api/v3",
	})

	return cfg
}
