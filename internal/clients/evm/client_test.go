package evm

import (
	"context"
	"github.com/ttdung/du/logger"
)

var (
	c   *EvmClient
	ctx context.Context
)

func init() {
	var err error

	c, err = NewEvmClient(DefaultConfig(), logger.NewZeroLogger(""))
	if err != nil {
		panic(err)
	}
	ctx = context.Background()
}
