package rpcchainlist

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/Helios-Chain-Labs/metrics"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
)

const (
	maxRespTime        = 15 * time.Second
	maxRespHeadersTime = 15 * time.Second
	maxRespBytes       = 10 * 1024 * 1024
)

var zeroPrice = float64(0)

type Config struct {
	BaseURL string
}

type RpcChainListFeed struct {
	client *http.Client
	config *Config

	logger  log.Logger
	svcTags metrics.Tags
}

// NewCoingeckoPriceFeed returns price puller for given symbol. The price will be pulled
// from endpoint and divided by scaleFactor. Symbol name (if reported by endpoint) must match.
func NewRpcChainListFeed() *RpcChainListFeed {
	return &RpcChainListFeed{
		client: &http.Client{
			Transport: &http.Transport{
				ResponseHeaderTimeout: maxRespHeadersTime,
			},
			Timeout: maxRespTime,
		},
		config: &Config{
			BaseURL: "https://chainlist.org/rpcs.json",
		},

		logger: log.WithFields(log.Fields{
			"svc":      "oracle",
			"provider": "rpcchainlist",
		}),
		svcTags: metrics.Tags{
			"provider": string("rpcchainlist"),
		},
	}
}

func (cp *RpcChainListFeed) QueryRpcs(chainId uint64) ([]string, error) {
	metrics.ReportFuncCall(cp.svcTags)
	doneFn := metrics.ReportFuncTiming(cp.svcTags)
	defer doneFn()

	u, err := url.ParseRequestURI(cp.config.BaseURL)
	if err != nil {
		metrics.ReportFuncError(cp.svcTags)
		cp.logger.WithError(err).Fatalln("failed to parse URL")
	}

	reqURL := u.String()
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		metrics.ReportFuncError(cp.svcTags)
		cp.logger.WithError(err).Fatalln("failed to create HTTP request")
	}

	resp, err := cp.client.Do(req)
	if err != nil {
		metrics.ReportFuncError(cp.svcTags)
		return nil, errors.Wrapf(err, "failed to fetch rpcs list from %s", reqURL)
	}

	respBody, err := ioutil.ReadAll(io.LimitReader(resp.Body, maxRespBytes))
	if err != nil {
		_ = resp.Body.Close()
		metrics.ReportFuncError(cp.svcTags)
		return nil, errors.Wrapf(err, "failed to read response body from %s", reqURL)
	}

	_ = resp.Body.Close()

	var f interface{}
	if err := json.Unmarshal(respBody, &f); err != nil {
		metrics.ReportFuncError(cp.svcTags)
		return nil, err
	}

	m, ok := f.([]interface{})
	if !ok {
		metrics.ReportFuncError(cp.svcTags)
		return nil, errors.Errorf("failed to cast response type: map[string]interface{}")
	}

	for _, chain := range m {
		chainMap := chain.(map[string]interface{})
		if chainMap["chainId"].(float64) == float64(chainId) {

			rpcs, ok := chainMap["rpc"].([]interface{})

			if !ok {
				return nil, errors.Errorf("failed to cast value type: map[string]interface{}")
			}

			rpcList := []string{}
			for _, rpc := range rpcs {
				rpcMap := rpc.(map[string]interface{})
				rpcList = append(rpcList, rpcMap["url"].(string))
			}

			return rpcList, nil
		}
	}
	return nil, errors.Errorf("failed to find chainId")
}
