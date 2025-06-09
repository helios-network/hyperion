package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/cmd/hyperion/queries"
	"github.com/Helios-Chain-Labs/hyperion/cmd/hyperion/static"
	globaltypes "github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/version"
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
		global.InitHeliosNetwork(0)

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		// global.StartRunnersAtStartUp(func(ctx context.Context, g *globaltypes.Global, chainId uint64) error {
		// 	return queries.RunHyperion(ctx, g, chainId)
		// })

		// API endpoints first
		apiRouter := router.PathPrefix("/api").Subrouter()
		apiRouter.HandleFunc("/query", injectGlobalMiddleware(handleQueryGet, global, ctx)).Methods("GET")
		apiRouter.HandleFunc("/query", injectGlobalMiddleware(handleQueryPost, global, ctx)).Methods("POST")
		apiRouter.HandleFunc("/version", handleVersion).Methods("GET")

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
	version := version.Version()
	sendSuccess(w, version, nil)
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

		// TODO: Implement actual password verification
		if params.Password != "your-secure-password" {
			sendError(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		sendSuccess(w, "Login successful", nil)
		return

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
		ctxGlobal := r.Context().Value(CtxKey).(context.Context)

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
