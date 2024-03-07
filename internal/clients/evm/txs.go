package evm

import (
	"context"
	internalCommon "github.com/ttdung/du/internal/common"
	"math/big"
	"time"
)

var (
	blockRetryInterval = 5 * time.Second
)

func (c *EvmClient) ListenToTxs(ctx context.Context, txResult chan interface{}, startBlk *big.Int) {
	var currentBlk *big.Int
	if startBlk != nil {
		currentBlk = new(big.Int).SetUint64(startBlk.Uint64())
	}
	for {
		select {
		case <-ctx.Done():
			c.log.Infof("ListenToTxs STOPPED")
			return

		default:
			head, err := c.LatestBlockHeight(ctx)
			if err != nil {
				c.log.Error("Unable to get latest block")
				time.Sleep(internalCommon.DefaultSleepTime)
				continue
			}
			if currentBlk == nil || currentBlk.Cmp(new(big.Int).SetUint64(0)) <= 0 {
				currentBlk = big.NewInt(head.Int64())
			}
			if head.Cmp(currentBlk) < 0 {
				time.Sleep(internalCommon.DefaultSleepTime)
				continue
			}

			c.log.Infof("================= BlockHeight: %v", currentBlk)

			txs, err := c.BlockTxsByHeight(ctx, currentBlk)
			if err != nil {
				c.log.Errorf("failed to get blockTxsByHeight(%v): %v", currentBlk.Uint64(), err)
				continue
			}
			//c.log.Infof("============== # of TX: %v", len(txs))
			for _, tx := range txs {
				c.log.Infof("============== TX hash: %v", tx.TxHash)
				txResult <- tx
			}

			if currentBlk.Uint64()%100 == 0 {
				c.log.Debugf("ListenToTxs finished block %v", currentBlk.Uint64())
			}
			currentBlk = currentBlk.Add(currentBlk, big.NewInt(1))
		}
	}
}
