package config

import (
	"encoding/json"
	"fmt"
	"github.com/ttdung/du/internal/clients"
	"github.com/ttdung/du/listener"
	"github.com/ttdung/du/logger"

	"io/ioutil"
)

type Config struct {
	Logger   logger.LoggerConfig     `json:"Logger"`
	Clients  clients.ClientsConfig   `json:"Clients"`
	Listener listener.ListenerConfig `json:"Listener"`
}

func DefaultConfig() Config {
	return Config{
		Logger:   logger.DefaultConfig(),
		Clients:  clients.DefaultConfig(),
		Listener: listener.DefaultConfig(),
	}
}

func (cfg Config) IsValid() (bool, error) {
	if _, err := cfg.Logger.IsValid(); err != nil {
		return false, fmt.Errorf("invalid LoggerConfig: %v", err)
	}

	if _, err := cfg.Clients.IsValid(); err != nil {
		return false, fmt.Errorf("invalid service config: %v", err)
	}

	return true, nil
}

// LoadConfigFromFile creates a new Config from the given file.
func LoadConfigFromFile(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var ret Config
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	if _, err = ret.IsValid(); err != nil {
		return nil, err
	}

	return &ret, nil
}

func SaveConfigToFile(cfg Config, filePath string) error {
	toBeWritten, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, toBeWritten, 0666)
}
