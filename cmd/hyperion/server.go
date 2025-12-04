package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/cmd/hyperion/queries"
	"github.com/Helios-Chain-Labs/hyperion/cmd/hyperion/static"
	globaltypes "github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/version"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/gorilla/mux"
	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
)

type contextKey string

const (
	GlobalKey contextKey = "global"
	CtxKey    contextKey = "ctx"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func injectGlobalMiddleware(next http.HandlerFunc, global *globaltypes.Global, ctxGlobal context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), GlobalKey, global)
		ctx = context.WithValue(ctx, CtxKey, ctxGlobal)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Skip auth check for login endpoint
		if r.URL.Query().Get("type") == "login" {
			next.ServeHTTP(w, r)
			return
		}

		// Get password from header
		providedPassword := r.Header.Get("X-Password")
		if providedPassword == "" {
			sendError(w, "Missing authentication", http.StatusUnauthorized)
			return
		}

		// Verify password
		storedPassword, err := storage.GetHyperionPassword()
		if err != nil {
			sendError(w, "Failed to verify authentication", http.StatusInternalServerError)
			return
		}

		if storedPassword == "" {
			sendError(w, "No password configured", http.StatusUnauthorized)
			return
		}

		if providedPassword != storedPassword {
			sendError(w, "Invalid authentication", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
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
		router.Use(loggingMiddleware)

		global := globaltypes.NewGlobal(&globaltypes.Config{
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
		heliosNetwork := global.GetHeliosNetwork()
		if heliosNetwork == nil {
			log.Fatal("helios network not initialized")
		}

		// ctx, cancelFn := context.WithCancel(context.Background())
		rootCtx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

		// never use rootStop, it will cause the server to hang when the context is cancelled
		// closer.Bind(rootStop)

		// API endpoints first
		apiRouter := router.PathPrefix("/api").Subrouter()
		apiRouter.HandleFunc("/query", authMiddleware(injectGlobalMiddleware(handleQueryGet, global, rootCtx))).Methods("GET")
		apiRouter.HandleFunc("/query", authMiddleware(injectGlobalMiddleware(handleQueryPost, global, rootCtx))).Methods("POST")
		apiRouter.HandleFunc("/version", handleVersion).Methods("GET")
		apiRouter.HandleFunc("/debug-goroutines", handleDebugGoroutines).Methods("GET")
		apiRouter.HandleFunc("/debug-goroutines-stats", handleDebugGoroutinesStats).Methods("GET")

		// Create file servers for both physical and embedded files
		// physicalFs := http.FileServer(http.Dir("static"))
		// embeddedFs := http.FileServer(static.GetIndex())
		indexFile := static.GetIndex()

		router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if requesting root path
			if r.URL.Path == "/" {
				// Try to serve physical index.html first
				if _, err := os.Stat("static/index.html"); err == nil {
					http.ServeFile(w, r, "static/index.html")
					return
				}
				// Fall back to embedded index.html
				r.URL.Path = "index.html"
				http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(indexFile))
				return
			}
		}))

		// Start runners at start up
		// go func() {
		// 	global.StartRunnersAtStartUp(func(ctx context.Context, g *globaltypes.Global, chainId uint64) error {
		// 		return queries.RunHyperion(ctx, g, chainId)
		// 	})
		// }()

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

func handleVersion(w http.ResponseWriter, r *http.Request) {
	// Get version info
	ver := version.Version()
	sendSuccess(w, ver, nil)
}

func handleDebugGoroutines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	pprof.Lookup("goroutine").WriteTo(w, 1)
}

func handleDebugGoroutinesStats(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	goroutineCount := runtime.NumGoroutine()

	// Get detailed goroutine info
	goroutineProfile := pprof.Lookup("goroutine")
	var buf bytes.Buffer
	goroutineProfile.WriteTo(&buf, 1)
	goroutineInfo := buf.String()

	// Count goroutines by function (basic analysis)
	goroutineStats := analyzeGoroutines(goroutineInfo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	stats := map[string]interface{}{
		"goroutine_count": goroutineCount,
		"memory_stats": map[string]interface{}{
			"heap_alloc":     m.HeapAlloc,
			"heap_sys":       m.HeapSys,
			"heap_idle":      m.HeapIdle,
			"heap_inuse":     m.HeapInuse,
			"total_alloc":    m.TotalAlloc,
			"gc_cycles":      m.NumGC,
			"gc_pause_total": m.PauseTotalNs,
		},
		"goroutine_analysis": goroutineStats,
		"timestamp":          time.Now().UTC().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(stats)
}

func handleQueryGet(w http.ResponseWriter, r *http.Request) {
	// Execute query-get command
	// Implementation details here
	global := r.Context().Value(GlobalKey).(*globaltypes.Global)
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
	case "get-list-tokens":
		chainId, err := strconv.ParseUint(query.Get("chain_id"), 10, 64)
		if err != nil {
			sendError(w, "Invalid chain_id", http.StatusBadRequest)
			return
		}
		page, err := strconv.ParseUint(query.Get("page"), 10, 64)
		if err != nil {
			sendError(w, "Invalid page", http.StatusBadRequest)
			return
		}
		size, err := strconv.ParseUint(query.Get("size"), 10, 64)
		if err != nil {
			sendError(w, "Invalid size", http.StatusBadRequest)
			return
		}

		tokens, err := queries.GetListTokens(r.Context(), global, chainId, page, size)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, tokens, nil)
		return
	case "get-list-rpcs":
		chainId, err := strconv.ParseUint(query.Get("chain_id"), 10, 64)
		if err != nil {
			sendError(w, "Invalid chain_id", http.StatusBadRequest)
			return
		}
		rpcs, err := queries.GetListRpcs(r.Context(), global, chainId)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, rpcs, nil)
		return
	case "get-list-transactions":
		page, err := strconv.ParseInt(query.Get("page"), 10, 64)
		if err != nil {
			sendError(w, "Invalid page", http.StatusBadRequest)
			return
		}
		size, err := strconv.ParseInt(query.Get("size"), 10, 64)
		if err != nil {
			sendError(w, "Invalid size", http.StatusBadRequest)
			return
		}
		transactions, err := queries.GetListTransactions(r.Context(), global, int(page), int(size))
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, transactions, nil)
		return
	case "get-wallet-address":
		address, err := queries.GetWalletAddress(r.Context(), global)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, address, nil)
		return
	case "get-list-outgoing-txs":
		chainId, err := strconv.ParseUint(query.Get("chain_id"), 10, 64)
		if err != nil {
			sendError(w, "Invalid chain_id", http.StatusBadRequest)
			return
		}
		outgoingTxs, err := queries.GetListOutgoingTxs(r.Context(), global, chainId)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, outgoingTxs, nil)
		return
	case "get-chain-settings":
		chainId, err := strconv.ParseUint(query.Get("chain_id"), 10, 64)
		if err != nil {
			sendError(w, "Invalid chain_id", http.StatusBadRequest)
			return
		}
		settings, err := queries.GetChainSettings(r.Context(), global, chainId)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, settings, nil)
		return
	case "get-validator":
		validator, err := queries.GetValidator(r.Context(), global)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, validator, nil)
		return
	case "get-stats":
		stats, err := queries.GetStats(r.Context(), global)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, stats, nil)
		return
	case "get-list-proposals":
		page, err := strconv.ParseInt(query.Get("page"), 10, 64)
		if err != nil {
			sendError(w, "Invalid page", http.StatusBadRequest)
			return
		}
		size, err := strconv.ParseInt(query.Get("size"), 10, 64)
		if err != nil {
			sendError(w, "Invalid size", http.StatusBadRequest)
			return
		}
		proposals, err := queries.GetListProposals(r.Context(), global, int(page), int(size))
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, proposals, nil)
		return
	}
	sendSuccess(w, "404", nil)
}

func handleQueryPost(w http.ResponseWriter, r *http.Request) {
	// Execute query-post command
	// Implementation details here
	global := r.Context().Value(GlobalKey).(*globaltypes.Global)
	query := r.URL.Query()
	queryType := query.Get("type")

	switch queryType {
	case "login":
		var params struct {
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		password, err := storage.GetHyperionPassword()
		if err != nil {
			sendError(w, "Failed to get password", http.StatusInternalServerError)
			return
		}
		if password == "" { // first time login
			storage.SetHyperionPassword(params.Password)
			sendSuccess(w, "Login successful", nil)
			return
		}
		if time.Since(global.LastTryAuthTime) < 10*time.Second {
			sendError(w, "Too many login attempts", http.StatusTooManyRequests)
			return
		}
		global.LastTryAuthTime = time.Now()
		if params.Password != password {
			sendError(w, "Invalid password", http.StatusUnauthorized)
			return
		}
		sendSuccess(w, "Login successful", nil)
		return

	case "add-new-chain":
		var params struct {
			ChainID uint64 `json:"chain_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.AddNewChain(r.Context(), global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return

	case "deploy-hyperion-contract":
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
	case "propose-hyperion-update":
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
		response, err := queries.ProposeHyperionUpdate(r.Context(), global, params.Title, params.Description, params.BridgeChainId, params.BridgeChainName, params.AverageCounterpartyBlockTime)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return

	case "propose-add-whitelisted-address":
		var params struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			ChainID     uint64 `json:"chain_id"`
			Address     string `json:"address"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.ProposeAddWhitelistedAddress(r.Context(), global, params.Title, params.Description, params.ChainID, params.Address)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return
	case "update-chain-logo":
		var params struct {
			ChainID uint64 `json:"chain_id"`
			Logo    string `json:"logo"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response := queries.UpdateChainLogo(r.Context(), global, params.ChainID, params.Logo)
		sendSuccess(w, response, nil)
		return
	case "add-whitelisted-address":
		var params struct {
			ChainID uint64 `json:"chain_id"`
			Address string `json:"address"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.AddWhitelistedAddress(r.Context(), global, params.ChainID, params.Address)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return
	case "upload-logo":
		var params struct {
			Logo string `json:"logo"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response := queries.UploadLogo(r.Context(), global, params.Logo)
		sendSuccess(w, response, nil)
		return
	case "pause-or-unpause-deposit":
		var params struct {
			ChainID uint64 `json:"chain_id"`
			Pause   bool   `json:"pause"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.PauseOrUnpauseDeposit(r.Context(), global, params.ChainID, params.Pause)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return
	case "pause-or-unpause-withdrawal":
		var params struct {
			ChainID uint64 `json:"chain_id"`
			Pause   bool   `json:"pause"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.PauseOrUnpauseWithdrawal(r.Context(), global, params.ChainID, params.Pause)
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
		rootCtx := r.Context().Value(CtxKey).(context.Context)

		err := queries.RunHyperion(rootCtx, global, params.ChainID)
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
		ctxGlobal := r.Context().Value(CtxKey).(context.Context)
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
		ctxGlobal := r.Context().Value(CtxKey).(context.Context)
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
		ctxGlobal := r.Context().Value(CtxKey).(context.Context)
		err := queries.UnRegisterHyperion(ctxGlobal, global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Hyperion unregistered successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "deploy-erc20":
		var params struct {
			ChainID  uint64 `json:"chain_id"`
			Denom    string `json:"denom"`
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Decimals uint8  `json:"decimals"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.DeployHeliosTokenToChain(r.Context(), global, params.ChainID, params.Denom, params.Name, params.Symbol, params.Decimals)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return
	case "add-rpcs":
		var params struct {
			ChainID   uint64 `json:"chain_id"`
			Rpcs      string `json:"rpcs"`
			IsPrimary bool   `json:"is_primary"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err := queries.AddRpcs(r.Context(), global, params.ChainID, params.Rpcs, params.IsPrimary)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "RPCs added successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "remove-rpcs":
		var params struct {
			ChainID uint64 `json:"chain_id"`
			Rpcs    string `json:"rpcs"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err := queries.RemoveRpcs(r.Context(), global, params.ChainID, params.Rpcs)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "RPCs removed successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "delete-chain":
		var params struct {
			ChainID uint64 `json:"chain_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err := queries.DeleteChain(r.Context(), global, params.ChainID)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Chain deleted successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "reset-helios-client":
		err := queries.ResetHeliosClient(r.Context(), global)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Helios client reset successfully", nil)
		return
	case "set-primary-rpc":
		var params struct {
			ChainID uint64 `json:"chain_id"`
			RpcUrl  string `json:"rpc_url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err := queries.SetPrimaryRpc(r.Context(), global, params.ChainID, params.RpcUrl)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Primary RPC set successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "update-chain-settings":
		var params struct {
			ChainID  uint64                 `json:"chain_id"`
			Settings map[string]interface{} `json:"settings"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err := queries.UpdateChainSettings(r.Context(), global, params.ChainID, params.Settings)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, "Chain settings updated successfully for chain "+strconv.FormatUint(params.ChainID, 10), nil)
		return
	case "claim-tokens-of-old-contract":
		var params struct {
			ChainID       uint64 `json:"chain_id"`
			Contract      string `json:"contract"`
			TokenContract string `json:"token_contract"`
			AmountInt     int64  `json:"amount_int"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response := queries.ClaimTokensOfOldContract(r.Context(), global, params.ChainID, params.Contract, params.TokenContract, params.AmountInt)
		sendSuccess(w, response, nil)
		return

	case "vote-for-proposal":
		var params struct {
			ProposalID uint64 `json:"proposal_id"`
			VoteOption string `json:"vote_option"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		var voteOption govtypes.VoteOption
		switch params.VoteOption {
		case "yes":
			voteOption = govtypes.OptionYes
		case "no":
			voteOption = govtypes.OptionNo
		case "abstain":
			voteOption = govtypes.OptionAbstain
		default:
			sendError(w, "Invalid vote option", http.StatusBadRequest)
			return
		}
		response, err := queries.VoteForProposal(r.Context(), global, params.ProposalID, voteOption)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return
	case "mint-token":
		var params struct {
			ChainID         uint64  `json:"chain_id"`
			TokenAddress    string  `json:"token_address"`
			Amount          float64 `json:"amount"`
			ReceiverAddress string  `json:"receiver_address"`
			Decimals        uint64  `json:"decimals"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		response, err := queries.MintToken(r.Context(), global, params.ChainID, params.TokenAddress, params.Decimals, params.Amount, params.ReceiverAddress)
		if err != nil {
			sendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendSuccess(w, response, nil)
		return
	}
	sendError(w, "Unknown query type", http.StatusBadRequest)
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

// analyzeGoroutines analyzes goroutine profile and returns statistics
func analyzeGoroutines(profile string) map[string]interface{} {
	stats := make(map[string]interface{})

	// Count goroutines by function name
	lines := strings.Split(profile, "\n")
	totalGoroutines := 0

	for _, line := range lines {
		if strings.Contains(line, "goroutine") && strings.Contains(line, "state") {
			totalGoroutines++
		}

		// Look for common patterns
		if strings.Contains(line, "net/http") {
			stats["http_goroutines"] = getCount(stats, "http_goroutines") + 1
		}
		if strings.Contains(line, "websocket") {
			stats["websocket_goroutines"] = getCount(stats, "websocket_goroutines") + 1
		}
		if strings.Contains(line, "rate_limiter") {
			stats["rate_limiter_goroutines"] = getCount(stats, "rate_limiter_goroutines") + 1
		}
		if strings.Contains(line, "compute_time") {
			stats["compute_time_goroutines"] = getCount(stats, "compute_time_goroutines") + 1
		}
		if strings.Contains(line, "method_tracker") {
			stats["method_tracker_goroutines"] = getCount(stats, "method_tracker_goroutines") + 1
		}
		if strings.Contains(line, "blocked") {
			stats["blocked_goroutines"] = getCount(stats, "blocked_goroutines") + 1
		}
	}

	stats["total_analyzed"] = totalGoroutines
	return stats
}

// getCount safely gets a count from stats map
func getCount(stats map[string]interface{}, key string) int {
	if val, ok := stats[key]; ok {
		if count, ok := val.(int); ok {
			return count
		}
	}
	return 0
}
