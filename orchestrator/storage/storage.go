package storage

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

func GetRpcsFromStorge(chainId uint64) ([]string, time.Duration, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, 0, err
	}

	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	rpcsDirPath := filepath.Join(dirPath, "rpcs")
	if _, err := os.Stat(rpcsDirPath); os.IsNotExist(err) {
		os.MkdirAll(rpcsDirPath, 0755)
	}

	joinPath := filepath.Join(rpcsDirPath, strconv.FormatUint(chainId, 10)+".json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("[]"), 0644)
	}

	baseFile, err := os.ReadFile(joinPath)

	if err != nil {
		baseFile = []byte("[]")
	}

	var baseFileArray []string
	json.Unmarshal(baseFile, &baseFileArray)

	stat, _ := os.Stat(joinPath)

	stat.ModTime()

	return baseFileArray, time.Since(stat.ModTime()), nil
}

func UpdateRpcsToStorge(chainId uint64, rpcs []string) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	rpcsDirPath := filepath.Join(dirPath, "rpcs")
	if _, err := os.Stat(rpcsDirPath); os.IsNotExist(err) {
		os.MkdirAll(rpcsDirPath, 0755)
	}

	joinPath := filepath.Join(rpcsDirPath, strconv.FormatUint(chainId, 10)+".json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("[]"), 0644)
	}

	jsonData, err := json.Marshal(rpcs)
	if err != nil {
		return err
	}

	os.WriteFile(joinPath, jsonData, 0644)

	return nil
}

func GetHyperionContractInfo(chainId uint64) (map[string]interface{}, error) {
	hyperions, err := GetMyHyperionsDeployedAddresses()
	if err != nil {
		return nil, err
	}
	for _, hyperion := range hyperions {
		fmt.Println("hyperion: ", hyperion)
		if uint64(hyperion["chainId"].(float64)) == chainId {
			return hyperion, nil
		}
	}
	return nil, fmt.Errorf("hyperion contract info not found")
}

func GetMyHyperionsDeployedAddresses() ([]map[string]interface{}, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	joinPath := filepath.Join(dirPath, "hyperions.json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("[]"), 0644)
	}

	baseFile, err := os.ReadFile(joinPath)

	if err != nil {
		baseFile = []byte("[]")
	}

	var baseFileArray []map[string]interface{}
	json.Unmarshal(baseFile, &baseFileArray)

	return baseFileArray, nil
}

func UpdateHyperionContractInfo(chainId uint64, contractAddress string, info map[string]interface{}) error {
	hyperions, err := GetMyHyperionsDeployedAddresses()
	if err != nil {
		return err
	}
	for _, hyperion := range hyperions {
		if uint64(hyperion["chainId"].(float64)) == chainId && hyperion["hyperionAddress"].(string) == contractAddress {
			for key, value := range info {
				hyperion[key] = value
			}
		}
	}
	return UpdateMyHyperionsDeployedAddresses(hyperions)
}

func UpdateMyHyperionsDeployedAddresses(hyperions []map[string]interface{}) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	joinPath := filepath.Join(dirPath, "hyperions.json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("[]"), 0644)
	}

	jsonData, err := json.Marshal(hyperions)
	if err != nil {
		return err
	}

	os.WriteFile(joinPath, jsonData, 0644)

	return nil
}

func AddOneNewHyperionDeployedAddress(hyperion map[string]interface{}) error {
	hyperions, err := GetMyHyperionsDeployedAddresses()
	if err != nil {
		return err
	}
	hyperions = append(hyperions, hyperion)

	return UpdateMyHyperionsDeployedAddresses(hyperions)
}

func GetRunners() ([]map[string]interface{}, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	joinPath := filepath.Join(dirPath, "runners.json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("[]"), 0644)
	}

	baseFile, err := os.ReadFile(joinPath)

	if err != nil {
		baseFile = []byte("[]")
	}

	var baseFileArray []map[string]interface{}
	json.Unmarshal(baseFile, &baseFileArray)

	return baseFileArray, nil
}

func SetRunner(chainId uint64) error {
	runners, err := GetRunners()
	if err != nil {
		return err
	}
	runners = append(runners, map[string]interface{}{
		"chainId": chainId,
	})
	return UpdateRunners(runners)
}

func RemoveRunner(chainId uint64) error {
	runners, err := GetRunners()
	if err != nil {
		return err
	}

	// Find the index of the runner to remove
	index := -1
	for i, runner := range runners {
		if uint64(runner["chainId"].(float64)) == chainId {
			index = i
			break
		}
	}

	// If found, remove it safely
	if index >= 0 && index < len(runners) {
		// Remove the element by copying the last element to the index position
		// and then truncating the slice
		runners[index] = runners[len(runners)-1]
		runners = runners[:len(runners)-1]
	}

	return UpdateRunners(runners)
}

func UpdateRunners(runners []map[string]interface{}) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	joinPath := filepath.Join(dirPath, "runners.json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("[]"), 0644)
	}

	jsonData, err := json.Marshal(runners)
	if err != nil {
		return err
	}

	os.WriteFile(joinPath, jsonData, 0644)

	return nil
}
