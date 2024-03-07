package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ttdung/du/logger"
	"math/big"
)

type EvmClient struct {
	*ethclient.Client
	GEthClient *gethclient.Client
	RPCClient  *rpc.Client
	log        logger.Logger
	chainID    *big.Int
}

// NewEvmClient creates a new EvmClient.
func NewEvmClient(cfg EvmClientConfig, log logger.Logger) (*EvmClient, error) {
	rpcClient, err := rpc.Dial(cfg.Endpoint)
	if err != nil {
		return nil, err
	}

	client := ethclient.NewClient(rpcClient)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &EvmClient{
		Client:     client,
		GEthClient: gethclient.New(rpcClient),
		RPCClient:  rpcClient,
		log:        log.WithPrefix("evm-client"),
		chainID:    chainID,
	}, nil
}

func (c *EvmClient) ChainID() *big.Int {
	return c.chainID
}
