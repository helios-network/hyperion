package rpcs

// func TestRpc(ctx context.Context, g *global.Global, chainId uint64, rpcUrl string) ([]Rpc, error) {

// 	ethKeyFromAddress, signerFn, personalSignFn, _ := keys.InitEthereumAccountsManagerWithRandomKey(chainId)

// 	ethNetwork, _ := ethereum.NewNetwork(common.HexToAddress("0x0000000000000000000000000000000000000000"), ethKeyFromAddress, signerFn, personalSignFn, ethereum.NetworkConfig{
// 		EthNodeRPCs:           []*hyperiontypes.Rpc{&hyperiontypes.Rpc{Url: rpcUrl, Reputation: 1, LastHeightUsed: 1}},
// 		GasPriceAdjustment:    g.GetConfig().EthGasPriceAdjustment,
// 		MaxGasPrice:           g.GetConfig().EthMaxGasPrice,
// 		PendingTxWaitDuration: g.GetConfig().PendingTxWaitDuration,
// 	})

// 	ethNetwork.TestRpcs(context.Background())

// 	testedRpcs := ethNetwork.GetRpcs()

// 	rpcsToSave := make([]string, 0)
// 	for _, rpc := range testedRpcs {
// 		rpcsToSave = append(rpcsToSave, rpc.Url)
// 	}

// 	storage.UpdateRpcsToStorge(chainId, rpcsToSave)

// 	rpcFinalList := make([]*hyperiontypes.Rpc, 0)
// 	for _, rpc := range rpcsToSave {
// 		rpcFinalList = append(rpcFinalList, &hyperiontypes.Rpc{
// 			Url: rpc,
// 		})
// 	}
// }
