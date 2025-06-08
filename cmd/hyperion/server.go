package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Helios-Chain-Labs/hyperion/cmd/hyperion/queries"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/version"
	"github.com/gorilla/mux"
	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func injectGlobalMiddleware(next http.HandlerFunc, global *global.Global, ctxGlobal context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "global", global)
		ctx = context.WithValue(ctx, "ctx", ctxGlobal)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func startServer(cmd *cli.Cmd) {
	cmd.Before = func() {
		initMetrics(cmd)
	}

	cfg := initConfig(cmd)

	cmd.Action = func() {
		// ensure a clean exit
		defer closer.Close()

		router := mux.NewRouter()

		global := global.NewGlobal(&global.Config{
			PrivateKey:            *cfg.heliosPrivKey,
			HeliosChainID:         *cfg.heliosChainID,
			HeliosGRPC:            *cfg.heliosGRPC,
			TendermintRPC:         *cfg.tendermintRPC,
			HeliosGasPrices:       *cfg.heliosGasPrices,
			HeliosGas:             *cfg.heliosGas,
			EthGasPriceAdjustment: *cfg.ethGasPriceAdjustment,
			EthMaxGasPrice:        *cfg.ethMaxGasPrice,
			PendingTxWaitDuration: *cfg.pendingTxWaitDuration,
		})
		global.InitHeliosNetwork(0)

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		log.Infof("Enabled logs: %s", *cfg.enabledLogs)

		// Query endpoints
		router.HandleFunc("/api/query", injectGlobalMiddleware(handleQueryGet, global, ctx)).Methods("GET")
		router.HandleFunc("/api/query", injectGlobalMiddleware(handleQueryPost, global, ctx)).Methods("POST")

		// Version endpoint
		router.HandleFunc("/api/version", handleVersion).Methods("GET")

		// Start server
		port := os.Getenv("PORT")
		if port == "" {
			port = "4040"
		}

		log.Infof("Starting server on port %s", port)
		if err := http.ListenAndServe(":"+port, router); err != nil {
			log.Fatal(err)
		}

		closer.Hold()
	}
}

func handleUnsetOrchestrator(w http.ResponseWriter, r *http.Request) {
	// Parse request body for parameters
	var params struct {
		ChainID uint64 `json:"chain_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Execute command
	cmd := cli.App("hyperion", "")
	cmd.Command("q query", "", queryCmdSubset)

	// Call unset-orchestrator with params
	// Implementation details here

	sendSuccess(w, "Orchestrator unset successfully", nil)
}

func handleSetOrchestrator(w http.ResponseWriter, r *http.Request) {
	var params struct {
		ChainID uint64 `json:"chain_id"`
		Address string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Execute set-orchestrator command
	// Implementation details here

	sendSuccess(w, "Orchestrator set successfully", nil)
}

// func handleListOperativeChains(w http.ResponseWriter, r *http.Request) {
// 	// Execute list-operative-chains command
// 	// Implementation details here

// 	global := r.Context().Value("global").(*global.Global)

// 	// global.TestRpcsAndGetRpcs(1, []string{})
// 	_, err := global.InitHeliosNetwork(1)
// 	if err != nil {
// 		sendError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	proposalId, err := global.CreateNewBlockchainProposal()
// 	if err != nil {
// 		sendError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println("Proposal ID:", proposalId)

// 	proposal, err := global.GetProposal(proposalId)
// 	if err != nil {
// 		sendError(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	decoded := map[string]interface{}{
// 		"status":           proposal.Status,
// 		"title":            proposal.Title,
// 		"proposer":         proposal.Proposer,
// 		"deposit_end_time": proposal.VotingEndTime,
// 	}

// 	// chains := []string{"chain1", "chain2"} // Example response
// 	sendSuccess(w, decoded, nil)
// }

func handleInitializeBlockchain(w http.ResponseWriter, r *http.Request) {
	var params struct {
		ChainID uint64 `json:"chain_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Execute initialize-blockchain command
	// Implementation details here

	sendSuccess(w, "Blockchain initialized successfully", nil)
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	// Get version info
	version := version.Version()
	sendSuccess(w, version, nil)
}

func handleDeployHyperion(w http.ResponseWriter, r *http.Request) {
	// Execute deploy-hyperion command
	// Implementation details here
	global := r.Context().Value("global").(*global.Global)

	address, _, success := global.DeployNewHyperionContract(11155111)
	if !success {
		sendError(w, "Failed to deploy Hyperion", http.StatusInternalServerError)
		return
	}

	fmt.Println("Hyperion deployed to:", address.String())

	sendSuccess(w, address.String(), nil)
}

func handleQueryGet(w http.ResponseWriter, r *http.Request) {
	// Execute query-get command
	// Implementation details here
	global := r.Context().Value("global").(*global.Global)
	query := r.URL.Query()
	queryType := query.Get("type")

	switch queryType {
	case "get-list-hyperions":
		hyperions, err := queries.GetListHyperions(r.Context(), global)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, hyperions, nil)
		return
	case "get-hyperion-info":
		infos, err := storage.GetHyperionContractInfo(11155111)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, infos, nil)
		return
	}
	sendSuccess(w, "404", nil)
}

func handleQueryPost(w http.ResponseWriter, r *http.Request) {
	// Execute query-post command
	// Implementation details here
	global := r.Context().Value("global").(*global.Global)
	query := r.URL.Query()
	queryType := query.Get("type")

	switch queryType {
	case "deploy-hyperion":
		var params struct {
			ChainID uint64 `json:"chain_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		hyperionContractInfo, err := queries.CreateHyperionContract(r.Context(), global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, hyperionContractInfo, nil)
		return
	case "propose-hyperion":
		var params struct {
			Title                        string `json:"title"`
			Description                  string `json:"description"`
			BridgeChainId                uint64 `json:"bridge_chain_id"`
			BridgeChainName              string `json:"bridge_chain_name"`
			AverageCounterpartyBlockTime uint64 `json:"average_counterparty_block_time"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.ProposeHyperion(r.Context(), global, params.Title, params.Description, params.BridgeChainId, params.BridgeChainName, params.AverageCounterpartyBlockTime)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return
	case "run-hyperion":
		var params struct {
			ChainID uint64 `json:"chain_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		ctxGlobal := r.Context().Value("ctx").(context.Context)

		err := queries.RunHyperion(ctxGlobal, global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Hyperion started successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "stop-hyperion":
		var params struct {
			ChainID uint64 `json:"chain_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		ctxGlobal := r.Context().Value("ctx").(context.Context)
		err := queries.StopHyperion(ctxGlobal, global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Hyperion stopped successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "register-hyperion":
		var params struct {
			ChainID uint64 `json:"chain_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		ctxGlobal := r.Context().Value("ctx").(context.Context)
		err := queries.RegisterHyperion(ctxGlobal, global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Hyperion registered successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "unregister-hyperion":
		var params struct {
			ChainID uint64 `json:"chain_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		ctxGlobal := r.Context().Value("ctx").(context.Context)
		err := queries.UnRegisterHyperion(ctxGlobal, global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Hyperion unregistered successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	}
	sendSuccess(w, "404", nil)
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   message,
	})
}

func sendSuccess(w http.ResponseWriter, data interface{}, metadata interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    data,
	})
}
