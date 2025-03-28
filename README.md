# Helios's Hyperion

Hyperion is a Go implementation of the Hyperion Orchestrator for the Helios Chain.

[Architecture Breakdown](./docs/ARCHITECTURE.md)

## Important

Important Commands:

* `hyperion orchestrator` starts the orchestrator main loop.
* `hyperion tx register-eth-key` is a special command to submit an Ethereum key that will be used to sign messages on behalf of your Validator

## Installation

Get yourself `Go 1.22+` at <https://golang.org/dl/> first, then:

```sh
go get github.com/Helios-Chain-Labs/hyperion/orchestrator/cmd/...
```

## hyperion

Hyperion is a companion executable for orchestrating a Hyperion validator.

### Configuration

Use CLI args, flags or create `.env` with environment variables

### Usage

```sh
$ hyperion --help

Usage: hyperion [OPTIONS] COMMAND [arg...]

Hyperion is a companion executable for orchestrating a Hyperion validator.

Options:
  -e, --env                The environment name this app runs in. Used for metrics and error reporting. (env $HYPERION_ENV) (default "local")
  -l, --log-level          Available levels: error, warn, info, debug. (env $HYPERION_LOG_LEVEL) (default "info")
      --svc-wait-timeout   Standard wait timeout for external services (e.g. Cosmos daemon GRPC connection) (env $HYPERION_SERVICE_WAIT_TIMEOUT) (default "1m")

Commands:
  orchestrator             Starts the orchestrator main loop.
  q, query                 Query commands that can get state info from Hyperion.
  version                  Print the version information and exit.

Run 'hyperion COMMAND --help' for more information on a command.      
```

## Commands

### hyperion orchestrator

```sh
$ hyperion orchestrator -h

Usage: hyperion orchestrator [OPTIONS]

Starts the orchestrator main loop.

Options:
      --cosmos-chain-id                  Specify Chain ID of the Cosmos network. (env $HYPERION_COSMOS_CHAIN_ID) (default "888")
      --cosmos-grpc                      Cosmos GRPC querying endpoint (env $HYPERION_COSMOS_GRPC) (default "tcp://localhost:9900")
      --tendermint-rpc                   Tendermint RPC endpoint (env $HYPERION_TENDERMINT_RPC) (default "http://localhost:26657")
      --cosmos-gas-prices                Specify Cosmos chain transaction fees as DecCoins gas prices (env $HYPERION_COSMOS_GAS_PRICES)
      --cosmos-keyring                   Specify Cosmos keyring backend (os|file|kwallet|pass|test) (env $HYPERION_COSMOS_KEYRING) (default "file")
      --cosmos-keyring-dir               Specify Cosmos keyring dir, if using file keyring. (env $HYPERION_COSMOS_KEYRING_DIR)
      --cosmos-keyring-app               Specify Cosmos keyring app name. (env $HYPERION_COSMOS_KEYRING_APP) (default "hyperion")
      --cosmos-from                      Specify the Cosmos validator key name or address. If specified, must exist in keyring, ledger or match the privkey. (env $HYPERION_COSMOS_FROM)
      --cosmos-from-passphrase           Specify keyring passphrase, otherwise Stdin will be used. (env $HYPERION_COSMOS_FROM_PASSPHRASE) (default "hyperion")
      --cosmos-pk                        Provide a raw Cosmos account private key of the validator in hex. USE FOR TESTING ONLY! (env $HYPERION_COSMOS_PK)
      --cosmos-use-ledger                Use the Cosmos app on hardware ledger to sign transactions. (env $HYPERION_COSMOS_USE_LEDGER)
      --eth-chain-id                     Specify Chain ID of the Ethereum network. (env $HYPERION_ETH_CHAIN_ID) (default 42)
      --eth-node-http                    Specify HTTP endpoint for an Ethereum node. (env $HYPERION_ETH_RPC) (default "http://localhost:1317")
      --eth-node-alchemy-ws              Specify websocket url for an Alchemy ethereum node. (env $HYPERION_ETH_ALCHEMY_WS)
      --eth_gas_price_adjustment         gas price adjustment for Ethereum transactions (env $HYPERION_ETH_GAS_PRICE_ADJUSTMENT) (default 1.3)
      --eth-keystore-dir                 Specify Ethereum keystore dir (Geth-format) prefix. (env $HYPERION_ETH_KEYSTORE_DIR)
      --eth-from                         Specify the from address. If specified, must exist in keystore, ledger or match the privkey. (env $HYPERION_ETH_FROM)
      --eth-passphrase                   Passphrase to unlock the private key from armor, if empty then stdin is used. (env $HYPERION_ETH_PASSPHRASE)
      --eth-pk                           Provide a raw Ethereum private key of the validator in hex. USE FOR TESTING ONLY! (env $HYPERION_ETH_PK)
      --eth-use-ledger                   Use the Ethereum app on hardware ledger to sign transactions. (env $HYPERION_ETH_USE_LEDGER)
      --relay_valsets                    If enabled, relayer will relay valsets to ethereum (env $HYPERION_RELAY_VALSETS)
      --relay_valset_offset_dur          If set, relayer will broadcast valsetUpdate only after relayValsetOffsetDur has passed from time of valsetUpdate creation (env $HYPERION_RELAY_VALSET_OFFSET_DUR) (default "5m")
      --relay_batches                    If enabled, relayer will relay batches to ethereum (env $HYPERION_RELAY_BATCHES)
      --relay_batch_offset_dur           If set, relayer will broadcast batches only after relayBatchOffsetDur has passed from time of batch creation (env $HYPERION_RELAY_BATCH_OFFSET_DUR) (default "5m")
      --relay_pending_tx_wait_duration   If set, relayer will broadcast pending batches/valsetupdate only after pendingTxWaitDuration has passed (env $HYPERION_RELAY_PENDING_TX_WAIT_DURATION) (default "20m")
      --min_batch_fee_usd                If set, batch request will create batches only if fee threshold exceeds (env $HYPERION_MIN_BATCH_FEE_USD) (default 23.3)
      --coingecko_api                    Specify HTTP endpoint for coingecko api. (env $HYPERION_COINGECKO_API) (default "https://api.coingecko.com/api/v3")
      --hyperion-id                      Specify Hyperion ID ecosystem. (env $HYPERION_ID)
```

## License

Apache 2.0

## Credit

Thanks to Injective Labs team who have worked on Peggo was a big inpirent project for working on Hyperion.
