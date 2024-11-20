#!/bin/bash

set -e

cd "${0%/*}" # cd in the script dir

PEGGY_ID="${PEGGY_ID:-0x696e6a6563746976652d70656767796964000000000000000000000000000000}"
POWER_THRESHOLD="${POWER_THRESHOLD:-1431655765}"
VALIDATOR_ADDRESSES="${VALIDATOR_ADDRESSES:-0xd561c412d17A1E16f0d6fbAc2F4eBDFD593Bb703}"
VALIDATOR_POWERS="${VALIDATOR_POWERS:-2147483647}"

if [[ ! -f .env ]]; then
	echo "Please create .env file, example is in .env.example"
	exit 1
fi

# peggy_impl_address=`etherman \
# 	--name Peggy \
# 	--source ../contracts/Peggy.sol \
# 	deploy`

# echo "Deployed Peggy implementation contract: $peggy_impl_address"
# echo -e "===\n"

# peggy_init_data=`etherman \
# 	--name Peggy \
# 	--source ../contracts/Peggy.sol \
# 	tx --await $peggy_impl_address initialize \
# 	$PEGGY_ID \
# 	$POWER_THRESHOLD \
# 	$VALIDATOR_ADDRESSES \
# 	$VALIDATOR_POWERS`

# echo "Using PEGGY_ID $PEGGY_ID"
# echo "Using POWER_THRESHOLD $POWER_THRESHOLD"
# echo "Using VALIDATOR_ADDRESSES $VALIDATOR_ADDRESSES"
# echo "Using VALIDATOR_POWERS $VALIDATOR_POWERS"
# echo -e "===\n"
# echo "Peggy Init data: $peggy_init_data"
# echo -e "===\n"

# echo "Peggy deployment done! Use $peggy_impl_address"

# deploy_erc20_log=`etherman \
# 	--name Peggy \
# 	--source ../contracts/Peggy.sol \
# 	logs $peggy_impl_address $peggy_init_data ValsetUpdatedEvent`

# echo $deploy_erc20_log

peggy_impl_address="0xdde3f7a7e6c038242F72F498ed0A4658Aa52Aa31"

# COSMOS_DENOM="${COSMOS_DENOM:-"helios"}"
# ERC20_NAME="${ERC20_NAME:-"HELIOS"}"
# ERC20_SYMBOL="${ERC20_SYMBOL:-"HELIOS"}"
# ERC20_DECIMALS="${ERC20_DECIMALS:-18}"
# deploy_erc20_txhash=`etherman \
# 	--name Peggy \
# 	--source ../contracts/Peggy.sol \
# 	tx --await $peggy_impl_address deployERC20 \
# 	$COSMOS_DENOM \
# 	$ERC20_NAME \
# 	$ERC20_SYMBOL \
# 	$ERC20_DECIMALS`

# echo "deployERC20 DONE"
echo $deploy_erc20_txhash

# proxy_admin_address=`etherman \
# 	--name ProxyAdmin \
# 	--source ../contracts/@openzeppelin/contracts/ProxyAdmin.sol \
# 	deploy`

# echo "Deployed ProxyAdmin contract: $proxy_admin_address"
# echo -e "===\n"

# peggy_proxy_address=`etherman \
# 	--name TransparentUpgradeableProxy \
# 	--source ../contracts/@openzeppelin/contracts/TransparentUpgradeableProxy.sol \
# 	deploy $peggy_impl_address $proxy_admin_address $peggy_init_data`

# echo "Deployed TransparentUpgradeableProxy for $peggy_impl_address (Peggy), with $proxy_admin_address (ProxyAdmin) as the admin"
# echo -e "===\n"

# echo "Peggy deployment done! Use $peggy_proxy_address"
