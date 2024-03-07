package clients

import (
	"context"
	"github.com/ttdung/du/internal/clients/common"
	"github.com/ttdung/du/internal/clients/evm"
	"github.com/ttdung/du/logger"
	"math/big"
)

type Client interface {
	ListenToTxs(ctx context.Context, resultChan chan interface{}, fromBlk *big.Int)
}

// NewClientsFromConfig creates new Client's from the given config.
func NewClientsFromConfig(cfg ClientsConfig, log logger.Logger) (map[string]Client, error) {
	ret := make(map[string]Client)
	if cfg.Evm.Enabled {
		evmClient, err := evm.NewEvmClient(cfg.Evm, log)
		if err != nil {
			return nil, err
		}
		ret[common.EvmClientName] = evmClient
	}

	return ret, nil
}
