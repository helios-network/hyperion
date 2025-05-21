package storage

import (
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
)

func UpdateFeesFile(feesTaken *big.Int, tokenContract string, cost *big.Int, txHash string, blockHeight uint64, chainId uint64) {

	homePath, err := os.UserHomeDir()
	if err != nil {
		return
	}

	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	joinPath := filepath.Join(dirPath, "fees.json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("[]"), 0644)
	}

	baseFile, err := os.ReadFile(joinPath)

	if err != nil {
		baseFile = []byte("[]")
	}

	var baseFileArray []map[string]interface{}
	json.Unmarshal(baseFile, &baseFileArray)

	feesObject := map[string]interface{}{
		"fees_taken":     feesTaken.String(),
		"token_contract": tokenContract,
		"cost":           cost.String(),
		"tx_hash":        txHash,
		"block_height":   blockHeight,
		"chain_id":       chainId,
	}
	baseFileArray = append(baseFileArray, feesObject)
	json, err := json.Marshal(baseFileArray)
	if err != nil {
		return
	}
	os.WriteFile(joinPath, json, 0644)
}
