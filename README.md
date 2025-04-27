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
      --svc-wait-timeout   Standard wait timeout for external services (e.g. Helios daemon GRPC connection) (env $HYPERION_SERVICE_WAIT_TIMEOUT) (default "1m")

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
      --helios-chain-id                  Specify Chain ID of the Helios network. (env $HYPERION_HELIOS_CHAIN_ID) (default "42000")
      --helios-grpc                      Helios GRPC querying endpoint (env $HYPERION_HELIOS_GRPC) (default "tcp://localhost:9090")
      --tendermint-rpc                   Tendermint RPC endpoint (env $HYPERION_TENDERMINT_RPC) (default "http://localhost:26657")
      --helios-gas-prices                Specify Helios chain transaction fees as DecCoins gas prices (env $HYPERION_HELIOS_GAS_PRICES)
      --helios-gas                       Specify Helios chain transaction gas hyperion will pay maximally (env $HYPERION_HELIOS_GAS)
      --helios-keyring                   Specify Helios keyring backend (os|file|pass|test|local) (env $HYPERION_HELIOS_KEYRING) (default "local")
      --helios-keyring-dir               Specify Helios keyring dir, if using file keyring. (env $HYPERION_HELIOS_KEYRING_DIR)
      --helios-keyring-app               Specify Helios keyring app name. (env $HYPERION_HELIOS_KEYRING_APP) (default "hyperion")
      --helios-from                      Specify Helios validator hex address. (env $HYPERION_HELIOS_FROM)
      --helios-from-passphrase           Specify keyring passphrase, otherwise Stdin will be used. (env $HYPERION_HELIOS_FROM_PASSPHRASE) (default "hyperion")
      --helios-pk                        Provide a raw Helios account private key of the validator in hex. (env $HYPERION_HELIOS_PK)
      --eth_gas_price_adjustment         gas price adjustment for Ethereum transactions (env $HYPERION_ETH_GAS_PRICE_ADJUSTMENT) (default 1.3)
      --eth-keystore-dir                 Specify Ethereum keystore dir (Geth-format) prefix. (env $HYPERION_ETH_KEYSTORE_DIR)
      --eth-from                         Specify the from address. If specified, must exist in keystore, ledger or match the privkey. (env $HYPERION_ETH_FROM)
      --eth-passphrase                   Passphrase to unlock the private key from armor, if empty then stdin is used. (env $HYPERION_ETH_PASSPHRASE)
      --eth-pk                           Provide a raw Ethereum private key of the validator in hex. USE FOR TESTING ONLY! (env $HYPERION_ETH_PK)
      --relay-valsets                    If enabled, relayer will relay valsets to ethereum (env $HYPERION_RELAY_VALSETS)
      --relay-valset-offset-dur          If set, relayer will broadcast valsetUpdate only after relayValsetOffsetDur has passed from time of valsetUpdate creation (env $HYPERION_RELAY_VALSET_OFFSET_DUR) (default "5m")
      --relay-batches                    If enabled, relayer will relay batches to ethereum (env $HYPERION_RELAY_BATCHES)
      --relay-batch-offset-dur           If set, relayer will broadcast batches only after relayBatchOffsetDur has passed from time of batch creation (env $HYPERION_RELAY_BATCH_OFFSET_DUR) (default "5m")
      --relay-pending-tx-wait-duration   If set, relayer will broadcast pending batches/valsetupdate only after pendingTxWaitDuration has passed (env $HYPERION_RELAY_PENDING_TX_WAIT_DURATION) (default "20m")
      --min-batch-fee-usd                If set, batch request will create batches only if fee threshold exceeds (env $HYPERION_MIN_BATCH_FEE_USD) (default 23.3)
      --coingecko-api                    Specify HTTP endpoint for coingecko api. (env $HYPERION_COINGECKO_API) (default "https://api.coingecko.com/api/v3")
```

## Setup keys

1 - Example of Keyring management with specifical directory

```sh
HYPERION_HELIOS_KEYRING="local"
HYPERION_HELIOS_KEYRING_DIR="/Users/x/.heliades"
HYPERION_HELIOS_KEYRING_APP="cosmos"
HYPERION_HELIOS_FROM="0x17267eB1FEC301848d4B5140eDDCFC48945427Ab"
HYPERION_HELIOS_FROM_PASSPHRASE=
HYPERION_HELIOS_PK=
```

2 - Example of Keyring Private Key

```sh
HYPERION_HELIOS_KEYRING=
HYPERION_HELIOS_KEYRING_DIR=
HYPERION_HELIOS_KEYRING_APP=
HYPERION_HELIOS_FROM=
HYPERION_HELIOS_FROM_PASSPHRASE=
HYPERION_HELIOS_PK="08373636333636366363...44344343"
```

## License

Apache 2.0

## Credit

Thanks to Injective Labs team who have worked on Peggo was a big inpirent project for working on Hyperion.
