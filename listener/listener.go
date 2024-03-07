package listener

import (
	"context"
	"github.com/ttdung/du/internal/clients"
	"github.com/ttdung/du/internal/common"
	"github.com/ttdung/du/logger"

	"math/big"
	"sync"
	"time"
)

type Listener struct {
	clients map[string]clients.Client
	log     logger.Logger
	mtx     *sync.Mutex
	cfg     ListenerConfig
}

// NewListener creates a new listener.
func NewListener(cfg ListenerConfig, clients map[string]clients.Client, log logger.Logger) (*Listener, error) {
	return &Listener{
		clients: clients,
		log:     log.WithPrefix("Listener"),
		mtx:     new(sync.Mutex),
		cfg:     cfg,
	}, nil
}

func (l *Listener) Start(ctx context.Context) {
	resultChan := make(chan interface{})

	var startBlock *big.Int
	switch l.cfg.StartBlock {
	case 0:
		startBlock = nil
	default:
		startBlock = big.NewInt(l.cfg.StartBlock)
	}

	for _, c := range l.clients {
		go c.ListenToTxs(ctx, resultChan, startBlock)
	}

	l.log.Infof("STARTED")
	for {
		select {
		case <-ctx.Done():
			l.log.Infof("STOPPED")
			return
		case msg := <-resultChan:
			if err, ok := msg.(error); ok {
				l.log.Errorf("new error received: %v", err)
				continue
			}

			l.log.Infof("================= New msg: %v", msg)
		default:
			time.Sleep(common.DefaultSleepTime)
		}
	}
}
