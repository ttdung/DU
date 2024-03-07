package clients

import (
	"fmt"
	"github.com/ttdung/du/internal/clients/evm"
)

type ClientsConfig struct {
	Evm evm.EvmClientConfig `json:"Evm"`
}

func DefaultConfig() ClientsConfig {
	return ClientsConfig{
		Evm: evm.DefaultConfig(),
	}
}

// IsValid checks if the current ClientsConfig is valid.
func (cfg ClientsConfig) IsValid() (bool, error) {
	if _, err := cfg.Evm.IsValid(); err != nil {
		return false, fmt.Errorf("invalid Evm: %v", err)
	}

	return true, nil
}
