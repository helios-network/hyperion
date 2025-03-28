#!/usr/bin/env bash

if ! command -v jq &> /dev/null
then
    echo "⚠️ jq command could not be found!"
    echo "jq is a lightweight and flexible command-line JSON processor."
    echo "Install it by checking https://stedolan.github.io/jq/download/"
    exit 1
fi


if [ -z "$1" ]
  then
    echo "Run the script with valid orchestrator helios address as argument"
    exit 1
fi

YOUR_ORCHESTRATOR_INJ_ADDRESS=$1

echo "1. Check pending batches to be confirmed"
echo "SLASHING_CONDITION - You will be slashed if  a batch request is not confirmed within 25000 blocks upon creation."
batch=$(curl -s https://lcd.helios.network/hyperion/v1/batch/last?address=${YOUR_ORCHESTRATOR_INJ_ADDRESS})
result=$(echo ${batch} | jq '.batch | length')
if [ ${result} -eq 0 ]; then
        echo "(O) No pending batches"
else
        echo "(X) result :"
        echo "${batch}"
fi

echo ""
echo "2. Check pending valsets to be confirmed"
echo "SLASHING_CONDITION - You will be slashed if  a batch request is not confirmed within 25000 blocks upon creation."
valsets=$(curl -s https://lcd.helios.network/hyperion/v1/valset/last?address=${YOUR_ORCHESTRATOR_INJ_ADDRESS})
result=$(echo ${valsets} | jq '.valsets | length')
if [ ${result} -eq 0 ]; then
        echo "(O) No Pending Valsets"
else
        echo "(X) result : "
        echo "${valsets}"
fi

echo ""
echo "3. Check latest event broadcasted by hyperion is upto date"
echo "SLASHING_CONDITION - You will be slashed if  you don't broadcast an event within 25000 blocks and it's broadcasted by majority of validators.  This is disabled for now."
lon=$(curl -s https://lcd.helios.network/hyperion/v1/module_state | jq '.state.last_observed_nonce')
lce=$(curl -s https://lcd.helios.network/hyperion/v1/oracle/event/${YOUR_ORCHESTRATOR_INJ_ADDRESS} | jq '.last_claim_event.ethereum_event_nonce')
if [ ${lon} == ${lce} ]; then
        echo "(O) your hyperion is upto date"
else
        echo "(X) check hyperion last_observed_nonce:${lon}, last_claim_event.ethereum_event_nonce:${lce}"
fi