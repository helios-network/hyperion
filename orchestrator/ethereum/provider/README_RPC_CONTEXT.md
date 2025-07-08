# RPC URL dans le Contexte

Cette fonctionnalité permet de récupérer l'URL du RPC utilisé lors des appels Ethereum directement depuis le contexte.

## Fonctionnalités

### 1. Ajout d'URL RPC au contexte

```go
// Ajouter une URL RPC au contexte
ctxWithRPC := WithRPCURL(ctx, "https://mainnet.infura.io/v3/YOUR-PROJECT-ID")
```

### 2. Récupération de l'URL RPC depuis le contexte

```go
// Récupérer l'URL RPC (retourne une chaîne vide si non trouvée)
rpcURL := GetCurrentRPCURL(ctx)

// Vérifier si l'URL RPC est disponible dans le contexte
if rpcURL, ok := GetRPCURLFromContext(ctx); ok {
    fmt.Printf("RPC utilisé: %s\n", rpcURL)
}
```

### 3. Sélection du meilleur RPC basé sur la réputation

```go
// Sélectionner le RPC avec la meilleure réputation
bestRpcURL := provider.SelectBestRatedRpcInRpcPool()
if bestRpcURL != "" {
    fmt.Printf("Meilleur RPC sélectionné: %s\n", bestRpcURL)
}
```

### 4. Gestion manuelle de la réputation des RPC

```go
// Pénaliser un RPC pour un échec
provider.PenalizeRpc("https://mainnet.infura.io/v3/YOUR-PROJECT-ID", 2)

// Féliciter un RPC pour un succès
provider.PraiseRpc("https://mainnet.infura.io/v3/YOUR-PROJECT-ID", 3)
```

## Utilisation

### Exemple basique

```go
func exampleUsage(ctx context.Context, provider EVMProvider) {
    // Appeler une méthode qui utilisera un RPC
    balance, url, err := provider.Balance(ctx, account)
    if err != nil {
        return
    }
    
    // Récupérer l'URL du RPC utilisé pour cet appel
    rpcURL := GetCurrentRPCURL(ctx)
    if rpcURL != "" {
        fmt.Printf("RPC utilisé: %s\n", rpcURL)
    }
}
```

### Exemple avec RPC personnalisé

```go
func exampleWithCustomRPC(ctx context.Context, provider EVMProvider, customRPC string) {
    // Ajouter une URL RPC personnalisée au contexte
    ctxWithRPC := WithRPCURL(ctx, customRPC)
    
    // Faire des appels avec le contexte contenant l'URL RPC
    balance, _, err := provider.Balance(ctxWithRPC, account)
    if err != nil {
        return
    }
    
    // L'URL RPC sera disponible dans le contexte
    rpcURL := GetCurrentRPCURL(ctxWithRPC)
    fmt.Printf("RPC personnalisé utilisé: %s\n", rpcURL)
}
```

### Utilisation dans les Loops (Relayer et Oracle)

#### Relayer Loop

```go
func (l *relayer) relay(ctx context.Context) error {
    // Sélectionner le meilleur RPC basé sur la réputation
    bestRpcURL := l.ethereum.SelectBestRatedRpcInRpcPool()
    if bestRpcURL != "" {
        l.Log().WithField("selected_rpc", bestRpcURL).Debug("Selected best rated RPC for relay")
        // Ajouter le meilleur RPC au contexte
        ctx = provider.WithRPCURL(ctx, bestRpcURL)
    }
    
    // Exécuter les opérations de relay avec le meilleur RPC
    // Toutes les méthodes du pool utiliseront automatiquement le RPC du contexte
    // ...
}
```

#### Oracle Loop

```go
func (l *oracle) observeEthEvents(ctx context.Context) error {
    // Sélectionner le meilleur RPC basé sur la réputation
    bestRpcURL := l.ethereum.SelectBestRatedRpcInRpcPool()
    if bestRpcURL != "" {
        l.Log().WithField("selected_rpc", bestRpcURL).Debug("Selected best rated RPC for oracle")
        // Ajouter le meilleur RPC au contexte
        ctx = provider.WithRPCURL(ctx, bestRpcURL)
    }
    
    // Exécuter les opérations d'oracle avec le meilleur RPC
    // Toutes les méthodes du pool utiliseront automatiquement le RPC du contexte
    // ...
}
```

#### Utilisation automatique du RPC du contexte

Une fois que le RPC est ajouté au contexte, toutes les méthodes du pool RPC l'utilisent automatiquement :

```go
// CallEthClientWithRetry vérifie automatiquement le contexte
err := pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
    // Cette opération utilisera le RPC spécifié dans le contexte
    return client.BalanceAt(ctx, account, nil)
})

// CallRpcClientWithRetry vérifie automatiquement le contexte
err := pool.CallRpcClientWithRetry(ctx, func(client *rpc.Client) error {
    // Cette opération utilisera le RPC spécifié dans le contexte
    return client.CallContext(ctx, &result, "eth_getBalance", account, "latest")
})
```

#### Gestion de la réputation dans les loops

##### Relayer Loop - Exemple d'utilisation

```go
// Dans relayBatchs - Gestion des succès et échecs
txHash, cost, err := l.ethereum.SendPreparedTx(ctx, txData)
if err != nil {
    // Pénaliser le RPC utilisé pour cet échec
    lastUsedRpc := l.ethereum.GetLastUsedRpc()
    if lastUsedRpc != "" {
        l.ethereum.PenalizeRpc(lastUsedRpc, 2) // Pénalité de 2 points
        l.Log().WithField("rpc", lastUsedRpc).Debug("Penalized RPC for failed transaction")
    }
    // ... gestion de l'erreur
} else {
    // Féliciter le RPC utilisé pour ce succès
    lastUsedRpc := l.ethereum.GetLastUsedRpc()
    if lastUsedRpc != "" {
        l.ethereum.PraiseRpc(lastUsedRpc, 3) // Récompense de 3 points
        l.Log().WithField("rpc", lastUsedRpc).Debug("Praised RPC for successful transaction")
    }
    // ... traitement du succès
}
```

##### Oracle Loop - Exemple d'utilisation

```go
// Dans getLatestEthHeight - Gestion des succès et échecs
h, err := l.ethereum.GetHeaderByNumber(ctx, nil)
if err != nil {
    // Pénaliser le RPC utilisé pour cet échec
    lastUsedRpc := l.ethereum.GetLastUsedRpc()
    if lastUsedRpc != "" {
        l.ethereum.PenalizeRpc(lastUsedRpc, 1) // Pénalité de 1 point
        l.Log().WithField("rpc", lastUsedRpc).Debug("Penalized RPC for failed header request")
    }
    return errors.Wrap(err, "failed to get latest ethereum header")
}

// Féliciter le RPC utilisé pour ce succès
lastUsedRpc := l.ethereum.GetLastUsedRpc()
if lastUsedRpc != "" {
    l.ethereum.PraiseRpc(lastUsedRpc, 1) // Récompense de 1 point
    l.Log().WithField("rpc", lastUsedRpc).Debug("Praised RPC for successful header request")
}
```

## Implémentation technique

### Clés de contexte

```go
type contextKey string

const (
    RPCURLKey contextKey = "rpc_url"
)
```

### Fonctions utilitaires

- `WithRPCURL(ctx context.Context, rpcURL string) context.Context` : Ajoute l'URL RPC au contexte
- `GetRPCURLFromContext(ctx context.Context) (string, bool)` : Récupère l'URL RPC avec un indicateur de succès
- `GetCurrentRPCURL(ctx context.Context) string` : Récupère l'URL RPC ou retourne une chaîne vide
- `SelectBestRatedRpcInRpcPool() string` : Sélectionne le RPC avec la meilleure réputation
- `PenalizeRpc(rpcUrl string, penalty uint64)` : Pénalise un RPC en réduisant sa réputation
- `PraiseRpc(rpcUrl string, praise uint64)` : Félicite un RPC en augmentant sa réputation

### Nouvelles méthodes du pool RPC

Le pool RPC a été étendu avec de nouvelles méthodes pour supporter l'utilisation de RPC spécifiques :

- `CallEthClientWithSpecificRPC(ctx context.Context, rpcURL string, operation func(*ethclient.Client) error) error` : Exécute une opération avec un RPC spécifique
- `CallRpcClientWithSpecificRPC(ctx context.Context, rpcURL string, operation func(*rpc.Client) error) error` : Exécute une opération RPC avec un RPC spécifique

### Logique de sélection automatique

Les méthodes existantes ont été modifiées pour vérifier automatiquement le contexte :

1. **CallEthClientWithRetry** : Vérifie si un RPC est spécifié dans le contexte
   - Si oui : Utilise `CallEthClientWithSpecificRPC`
   - Sinon : Utilise la sélection aléatoire habituelle

2. **CallRpcClientWithRetry** : Vérifie si un RPC est spécifié dans le contexte
   - Si oui : Utilise `CallRpcClientWithSpecificRPC`
   - Sinon : Utilise la sélection round-robin habituelle

## Avantages

1. **Traçabilité** : Possibilité de savoir quel RPC a été utilisé pour chaque appel
2. **Debugging** : Facilite le débogage en cas de problème avec un RPC spécifique
3. **Monitoring** : Permet de surveiller l'utilisation des différents RPC
4. **Flexibilité** : Possibilité de forcer l'utilisation d'un RPC spécifique via le contexte
5. **Optimisation** : Sélection automatique du meilleur RPC basé sur la réputation
6. **Fiabilité** : Améliore la fiabilité des loops en utilisant les RPC les plus performants

## Notes importantes

- L'URL RPC est automatiquement ajoutée au contexte lors des appels via le pool de RPC
- Le contexte avec l'URL RPC est utilisé en interne pour tous les appels Ethereum
- Cette fonctionnalité est rétrocompatible et n'affecte pas le comportement existant
- La sélection du meilleur RPC est basée sur un système de réputation qui s'ajuste automatiquement
- Les loops relayer et oracle utilisent maintenant automatiquement le meilleur RPC disponible 