package hyperion

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (s *hyperionContract) EmergencyPause(ctx context.Context) (*common.Hash, error) {
	tx, err := s.ethHyperion.EmergencyPause(&bind.TransactOpts{
		From:   s.FromAddress(),
		Signer: s.signerFn,
	})
	if err != nil {
		return nil, err
	}
	hash := tx.Hash()
	return &hash, nil
}

func (s *hyperionContract) EmergencyUnpause(ctx context.Context) (*common.Hash, error) {
	tx, err := s.ethHyperion.EmergencyUnpause(&bind.TransactOpts{
		From:   s.FromAddress(),
		Signer: s.signerFn,
	})
	if err != nil {
		return nil, err
	}
	hash := tx.Hash()
	return &hash, nil
}

func (s *hyperionContract) PauseOrUnpauseDeposit(ctx context.Context, pause bool) (*common.Hash, error) {
	if pause {
		return s.EmergencyPause(ctx)
	} else {
		return s.EmergencyUnpause(ctx)
	}
}

func (s *hyperionContract) IsDepositPaused(ctx context.Context) (bool, error) {
	paused, err := s.ethHyperion.Paused(&bind.CallOpts{
		From: s.FromAddress(),
	})
	if err != nil {
		return false, err
	}
	return paused, nil
}
