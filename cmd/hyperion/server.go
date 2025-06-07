package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

func injectGlobalMiddleware(next http.HandlerFunc, global *Global) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "global", global)
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

		global := NewGlobal(&cfg)

		log.Infof("Enabled logs: %s", *cfg.enabledLogs)

		// Query endpoints
		router.HandleFunc("/api/query/unset-orchestrator", injectGlobalMiddleware(handleUnsetOrchestrator, global)).Methods("POST")
		router.HandleFunc("/api/query/set-orchestrator", injectGlobalMiddleware(handleSetOrchestrator, global)).Methods("POST")
		router.HandleFunc("/api/query/list-operative-chains", injectGlobalMiddleware(handleListOperativeChains, global)).Methods("GET")
		router.HandleFunc("/api/query/initialize-blockchain", injectGlobalMiddleware(handleInitializeBlockchain, global)).Methods("POST")
		router.HandleFunc("/api/query/deploy-hyperion", injectGlobalMiddleware(handleDeployHyperion, global)).Methods("GET")

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

func handleListOperativeChains(w http.ResponseWriter, r *http.Request) {
	// Execute list-operative-chains command
	// Implementation details here

	global := r.Context().Value("global").(*Global)

	// global.TestRpcsAndGetRpcs(1, []string{})
	_, err := global.InitHeliosNetwork(1)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	proposalId, err := global.CreateNewBlockchainProposal()
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Proposal ID:", proposalId)

	proposal, err := global.GetProposal(proposalId)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	decoded := map[string]interface{}{
		"status":           proposal.Status,
		"title":            proposal.Title,
		"proposer":         proposal.Proposer,
		"deposit_end_time": proposal.VotingEndTime,
	}

	// chains := []string{"chain1", "chain2"} // Example response
	sendSuccess(w, decoded, nil)
}

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
	global := r.Context().Value("global").(*Global)

	address, success := global.DeployNewHyperionContract(11155111)
	if !success {
		sendError(w, "Failed to deploy Hyperion", http.StatusInternalServerError)
		return
	}

	fmt.Println("Hyperion deployed to:", address.String())

	sendSuccess(w, address.String(), nil)
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
