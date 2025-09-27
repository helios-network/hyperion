package storage

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcs"
)

func GetFeesFile() ([]map[string]interface{}, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	var baseFileArray []map[string]interface{}
	json.Unmarshal(baseFile, &baseFileArray)

	for i := len(baseFileArray)/2 - 1; i >= 0; i-- {
		opp := len(baseFileArray) - 1 - i
		baseFileArray[i], baseFileArray[opp] = baseFileArray[opp], baseFileArray[i]
	}

	return baseFileArray, nil
}

func UpdateFeesFile(feesTaken *big.Int, tokenContract string, cost *big.Int, txHash string, blockHeight uint64, chainId uint64, txType string) {

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
		"tx_type":        txType,
		"chain_id":       chainId,
		"timestamp":      time.Now().Unix(),
	}
	baseFileArray = append(baseFileArray, feesObject)
	json, err := json.Marshal(baseFileArray)
	if err != nil {
		return
	}
	os.WriteFile(joinPath, json, 0644)
}

func GetRpcsFromStorge(chainId uint64) ([]*rpcs.Rpc, time.Duration, error) {
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

	var baseFileArray []*rpcs.Rpc
	json.Unmarshal(baseFile, &baseFileArray)

	stat, _ := os.Stat(joinPath)

	stat.ModTime()

	baseFileArray = OrderRpcsByPrimaryFirst(baseFileArray)

	return baseFileArray, time.Since(stat.ModTime()), nil
}

func AddRpcToStorge(chainId uint64, rpc *rpcs.Rpc) error {
	rpcsList, _, err := GetRpcsFromStorge(chainId)
	if err != nil {
		return err
	}
	alreadyExists := false
	for _, r := range rpcsList {
		if r.Url == rpc.Url {
			alreadyExists = true
			break
		}
	}
	if alreadyExists {
		return nil
	}
	rpcsList = append(rpcsList, rpc)
	return UpdateRpcsToStorge(chainId, rpcsList)
}

func RemoveRpcFromStorge(chainId uint64, rpc *rpcs.Rpc) error {
	rpcsList, _, err := GetRpcsFromStorge(chainId)
	if err != nil {
		return err
	}
	newRpcs := make([]*rpcs.Rpc, 0)
	for _, r := range rpcsList {
		if r.Url != rpc.Url {
			newRpcs = append(newRpcs, r)
		}
	}
	return UpdateRpcsToStorge(chainId, newRpcs)
}

func OrderRpcsByPrimaryFirst(rpcsList []*rpcs.Rpc) []*rpcs.Rpc {
	primaryRpcs := make([]*rpcs.Rpc, 0)
	secondaryRpcs := make([]*rpcs.Rpc, 0)
	for _, r := range rpcsList {
		if r.IsPrimary {
			primaryRpcs = append(primaryRpcs, r)
		} else {
			secondaryRpcs = append(secondaryRpcs, r)
		}
	}
	return append(primaryRpcs, secondaryRpcs...)
}

func UpdateRpcsToStorge(chainId uint64, rpcsList []*rpcs.Rpc) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	rpcsList = OrderRpcsByPrimaryFirst(rpcsList)

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

	jsonData, err := json.Marshal(rpcsList)
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
		if uint64(hyperion["chainId"].(float64)) == chainId && (hyperion["hyperionAddress"].(string) == contractAddress || hyperion["hyperionAddress"].(string) == "0x0000000000000000000000000000000000000000") {
			for key, value := range info {
				hyperion[key] = value
			}
		}
	}
	return UpdateMyHyperionsDeployedAddresses(hyperions)
}

func RemoveHyperionContractInfo(chainId uint64, contractAddress string) error {
	hyperions, err := GetMyHyperionsDeployedAddresses()
	if err != nil {
		return err
	}
	index := -1
	for i, hyperion := range hyperions {
		if uint64(hyperion["chainId"].(float64)) == chainId && hyperion["hyperionAddress"].(string) == contractAddress {
			index = i
			break
		}
	}
	if index >= 0 && index < len(hyperions) {
		hyperions[index] = hyperions[len(hyperions)-1]
		hyperions = hyperions[:len(hyperions)-1]
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

	alreadyExists := false
	for _, runner := range runners {
		if uint64(runner["chainId"].(float64)) == chainId {
			alreadyExists = true
			break
		}
	}
	if alreadyExists {
		RemoveRunner(chainId)
	}
	runners, err = GetRunners()
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

	newRunners := make([]map[string]interface{}, 0)
	for _, runner := range runners {
		if uint64(runner["chainId"].(float64)) == chainId {
			continue
		}
		newRunners = append(newRunners, runner)
	}

	return UpdateRunners(newRunners)
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

func SetHyperionPassword(password string) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}
	joinPath := filepath.Join(dirPath, "password.txt")
	os.WriteFile(joinPath, []byte(password), 0644)
	return nil
}

func GetHyperionPassword() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}
	joinPath := filepath.Join(dirPath, "password.txt")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte(""), 0644)
	}
	baseFile, err := os.ReadFile(joinPath)
	if err != nil {
		return "", err
	}
	return string(baseFile), nil
}

func SetChainSettings(chainId uint64, settings map[string]interface{}) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}
	joinPath := filepath.Join(dirPath, "chain_settings.json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("{}"), 0644)
	}
	baseFile, err := os.ReadFile(joinPath)
	if err != nil {
		return err
	}
	var baseFileMap map[string]interface{}
	json.Unmarshal(baseFile, &baseFileMap)

	baseFileMap[strconv.FormatUint(chainId, 10)] = settings
	jsonData, err := json.Marshal(baseFileMap)
	if err != nil {
		return err
	}
	os.WriteFile(joinPath, jsonData, 0644)
	return nil
}

func GetChainSettings(chainId uint64, defaultSettings map[string]interface{}) (map[string]interface{}, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dirPath := filepath.Join(homePath, ".heliades", "hyperion")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}
	joinPath := filepath.Join(dirPath, "chain_settings.json")
	if _, err := os.Stat(joinPath); os.IsNotExist(err) {
		os.WriteFile(joinPath, []byte("{}"), 0644)
	}
	baseFile, err := os.ReadFile(joinPath)
	if err != nil {
		return nil, err
	}
	var baseFileMap map[string]interface{}
	json.Unmarshal(baseFile, &baseFileMap)
	if _, ok := baseFileMap[strconv.FormatUint(chainId, 10)]; !ok {
		return defaultSettings, nil
	}
	return baseFileMap[strconv.FormatUint(chainId, 10)].(map[string]interface{}), nil
}
