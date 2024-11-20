#!/bin/bash

set -e

cd "${0%/*}" # cd in the script dir

PEGGY_ADDRESS="0x3F567E87c156D731E05DaEebeBe82E30BaD7c4b5"

if [[ ! -f .env ]]; then
	echo "Please create .env file, example is in .env.example"
	exit 1
fi

if [[ "$PEGGY_ADDRESS" == "" ]]; then
	echo "Please set \$PEGGY_ADDRESS variable to a deployed Peggy instance"
	exit 1
fi

deploy_erc20_txhash=`etherman \
	--name Peggy \
	--source ../contracts/Peggy.sol \
	tx $PEGGY_ADDRESS deployERC20 helios HELIOS HELIOS 18`

# echo "deployERC20 DONE"
echo $deploy_erc20_txhash

# deploy_erc20_log=`etherman \
# 	--name Peggy \
# 	--source ../contracts/Peggy.sol \
# 	logs $PEGGY_ADDRESS 0x5e82ef824cc0641c9ef432dd606958c6b5ffb7a5815059a61bc710faad01ef42 ERC20DeployedEvent`

# echo $deploy_erc20_log
# erc20_token_address=`jq -r '..|._tokenContract?' <<< $deploy_erc20_log`

# echo "Deployed Cosmos ERC20 INJ Contract $erc20_token_address"
