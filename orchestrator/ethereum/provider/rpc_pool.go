package provider

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// ErrNoClientsAvailable is returned when no RPC clients are available in the pool
	ErrNoClientsAvailable = errors.New("no RPC clients available in the pool")
	// ErrRequestTimeout is returned when all requests time out
	ErrRequestTimeout = errors.New("all requests timed out")
	// ErrAllAttemptsFailed is returned when all retry attempts fail
	ErrAllAttemptsFailed = errors.New("all retry attempts failed")
)

type RPCReputation struct {
	rpcUrl     string
	reputation uint64
}

// EVMProviders manages multiple Ethereum RPC and ethclient connections
type EVMProviders struct {
	rcs         []*rpc.Client
	ethClients  []*ethclient.Client
	urls        []string
	reputations map[string]*RPCReputation
	currentIdx  uint32
	maxRetries  int
	timeout     time.Duration
	mu          sync.RWMutex
}

// NewEVMProviders creates a new EVMProviders instance with the given RPC URLs
func NewEVMProviders(rpcs []*hyperiontypes.Rpc) *EVMProviders {
	return NewEVMProvidersWithOptions(rpcs, 3, 10*time.Second)
}

// NewEVMProvidersWithOptions creates a new EVMProviders with custom retry and timeout settings
func NewEVMProvidersWithOptions(rpcs []*hyperiontypes.Rpc, maxRetries int, timeout time.Duration) *EVMProviders {
	// Loop and create rpc clients
	var rcs []*rpc.Client
	var ethClients []*ethclient.Client
	var validUrls []string
	reputations := make(map[string]*RPCReputation)
	for _, rpcReputation := range rpcs {
		client, err := rpc.Dial(rpcReputation.Url)
		if err != nil {
			// Skip failed connections but don't panic
			continue
		}
		ethClient := ethclient.NewClient(client)
		rcs = append(rcs, client)
		ethClients = append(ethClients, ethClient)
		validUrls = append(validUrls, rpcReputation.Url)
		reputations[rpcReputation.Url] = &RPCReputation{
			rpcUrl:     rpcReputation.Url,
			reputation: rpcReputation.Reputation,
		}
	}

	if len(rcs) == 0 {
		// If no connections could be established, panic
		panic("failed to connect to any RPC endpoints")
	}

	return &EVMProviders{
		rcs:         rcs,
		ethClients:  ethClients,
		urls:        validUrls,
		maxRetries:  maxRetries,
		timeout:     timeout,
		reputations: reputations,
	}
}

// getNextClient returns the next client using round-robin selection
func (p *EVMProviders) getNextClient() (*ethclient.Client, *rpc.Client) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if len(p.ethClients) == 0 {
		return nil, nil
	}

	// Atomically increment and wrap around
	idx := atomic.AddUint32(&p.currentIdx, 1) % uint32(len(p.ethClients))
	return p.ethClients[idx], p.rcs[idx]
}

// getRandomClient returns a random client
func (p *EVMProviders) getRandomClient() (*ethclient.Client, *rpc.Client, string) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if len(p.ethClients) == 0 {
		return nil, nil, ""
	}

	idx := rand.Intn(len(p.ethClients))
	return p.ethClients[idx], p.rcs[idx], p.urls[idx]
}

func (p *EVMProviders) classifyRpcUrl(rpcUrl string, failed bool) {
	if !failed {
		p.reputations[rpcUrl].reputation += 1
	} else if p.reputations[rpcUrl].reputation > 0 {
		p.reputations[rpcUrl].reputation -= 1
	}
	fmt.Println("rpcUrl: ", rpcUrl, "reputation: ", p.reputations[rpcUrl].reputation)
}

// CallEthClientWithRetry executes an operation with the ethclient with retry logic
func (p *EVMProviders) CallEthClientWithRetry(ctx context.Context, operation func(*ethclient.Client) error) error {
	p.mu.RLock()
	clientCount := len(p.ethClients)
	p.mu.RUnlock()

	if clientCount == 0 {
		return ErrNoClientsAvailable
	}

	var lastErr error
	for attempt := 0; attempt < p.maxRetries; attempt++ {
		// Get next client using round-robin
		ethClient, _, rpcUrl := p.getRandomClient()
		if ethClient == nil {
			return ErrNoClientsAvailable
		}

		// Create a context with timeout
		timeoutCtx, cancel := context.WithTimeout(ctx, p.timeout)

		// Execute the operation
		errCh := make(chan error, 1)
		go func() {
			errCh <- operation(ethClient)
		}()

		// Wait for either the result or timeout
		select {
		case err := <-errCh:
			cancel()
			if err == nil {
				fmt.Println("SUCCESSrpcUrl: ", rpcUrl)
				p.classifyRpcUrl(rpcUrl, false)
				return nil // Success
			}
			lastErr = err
			fmt.Println("ERRORrpcUrl: ", rpcUrl)
			p.classifyRpcUrl(rpcUrl, true)
			// Continue to next attempt
		case <-timeoutCtx.Done():
			cancel()
			lastErr = timeoutCtx.Err()
			fmt.Println("TIMEOUTrpcUrl: ", rpcUrl)
			p.classifyRpcUrl(rpcUrl, true)
			// Continue to next attempt
		}
	}

	if lastErr != nil {
		return lastErr
	}
	return ErrAllAttemptsFailed
}

// CallRpcClientWithRetry executes an operation with the rpc.Client with retry logic
func (p *EVMProviders) CallRpcClientWithRetry(ctx context.Context, operation func(*rpc.Client) error) error {
	p.mu.RLock()
	clientCount := len(p.rcs)
	p.mu.RUnlock()

	if clientCount == 0 {
		return ErrNoClientsAvailable
	}

	var lastErr error
	for attempt := 0; attempt < p.maxRetries; attempt++ {
		// Get next client using round-robin
		_, rpcClient := p.getNextClient()
		if rpcClient == nil {
			return ErrNoClientsAvailable
		}

		// Create a context with timeout
		timeoutCtx, cancel := context.WithTimeout(ctx, p.timeout)

		// Execute the operation
		errCh := make(chan error, 1)
		go func() {
			errCh <- operation(rpcClient)
		}()

		// Wait for either the result or timeout
		select {
		case err := <-errCh:
			cancel()
			if err == nil {
				return nil // Success
			}
			lastErr = err
			// Continue to next attempt
		case <-timeoutCtx.Done():
			cancel()
			lastErr = timeoutCtx.Err()
			// Continue to next attempt
		}
	}

	if lastErr != nil {
		return lastErr
	}
	return ErrAllAttemptsFailed
}

// Close closes all clients
func (p *EVMProviders) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, client := range p.rcs {
		client.Close()
	}
	p.rcs = nil
	p.ethClients = nil
}

// AddEndpoint adds a new RPC endpoint to the provider
func (p *EVMProviders) AddEndpoint(rpcUrl string) error {
	client, err := rpc.Dial(rpcUrl)
	if err != nil {
		return err
	}

	ethClient := ethclient.NewClient(client)

	p.mu.Lock()
	defer p.mu.Unlock()

	p.rcs = append(p.rcs, client)
	p.ethClients = append(p.ethClients, ethClient)
	p.urls = append(p.urls, rpcUrl)

	return nil
}

// GetClientCount returns the number of available clients
func (p *EVMProviders) GetClientCount() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.ethClients)
}
