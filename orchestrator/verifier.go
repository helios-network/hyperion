package orchestrator

import (
	"context"
	"strings"
)

func VerifyTxError(ctx context.Context, err string, orchestrator *Orchestrator) (bool, error) {
	if strings.Contains(err, "account sequence mismatch") {
		orchestrator.HyperionState.ErrorStatus = "Error: Check Your Node - Maybe Jail or unsync"
		orchestrator.ResetHeliosClient()
	}
	return true, nil
}
