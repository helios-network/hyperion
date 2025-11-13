package ethereum

import (
	"context"
	"math/big"
	"sync"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type INetworkCached interface {
	GetHeaderByNumber(ctx context.Context, number *big.Int) (*gethtypes.Header, error)
}

// cacheEntry représente une entrée dans le cache avec son timestamp d'expiration
type cacheEntry struct {
	header    *gethtypes.Header
	expiresAt time.Time
}

// INetworkCachedImpl implémente INetworkCached avec un système de cache
type INetworkCachedImpl struct {
	network Network

	// Cache pour GetHeaderByNumber
	headerCache    map[string]*cacheEntry
	headerCacheMux sync.RWMutex
	headerCacheTTL time.Duration
}

// GetHeaderByNumber récupère un header de bloc avec cache
func (n *INetworkCachedImpl) GetHeaderByNumber(ctx context.Context, number *big.Int) (*gethtypes.Header, error) {
	// Créer une clé de cache basée sur le numéro de bloc
	cacheKey := ""
	if number != nil {
		cacheKey = number.String()
	} else {
		cacheKey = "latest"
	}

	// Vérifier le cache
	n.headerCacheMux.RLock()
	if entry, exists := n.headerCache[cacheKey]; exists {
		// Vérifier si l'entrée n'est pas expirée
		if time.Now().Before(entry.expiresAt) {
			n.headerCacheMux.RUnlock()
			return entry.header, nil
		}
	}
	n.headerCacheMux.RUnlock()

	// Cache miss ou expiré, faire l'appel réseau
	header, err := n.network.GetHeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	// Stocker dans le cache
	n.headerCacheMux.Lock()
	n.headerCache[cacheKey] = &cacheEntry{
		header:    header,
		expiresAt: time.Now().Add(n.headerCacheTTL),
	}
	n.headerCacheMux.Unlock()

	return header, nil
}

// NewINetworkCached crée une nouvelle instance avec cache configuré
func NewINetworkCached(network Network) INetworkCached {
	return &INetworkCachedImpl{
		network:        network,
		headerCache:    make(map[string]*cacheEntry),
		headerCacheTTL: 2 * time.Second,
	}
}

// NewINetworkCachedWithTTL crée une instance avec des TTL personnalisés
func NewINetworkCachedWithTTL(network Network, headerTTL time.Duration) INetworkCached {
	return &INetworkCachedImpl{
		network:        network,
		headerCache:    make(map[string]*cacheEntry),
		headerCacheTTL: headerTTL,
	}
}
