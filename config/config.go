package config

import (
	"encoding/json"
	"fmt"
	"github.com/ttdung/du/internal/clients"

	"io/ioutil"
)

type Config struct {
	Logger  logger.LoggerConfig   `json:"Logger"`
	Clients clients.ClientsConfig `json:"Clients"`
}

func DefaultConfig() Config {
	return Config{
		Logger: logger.DefaultConfig(),
		//Webhooks:   webhook.DefaultConfig(),
		Clients: clients.DefaultConfig(),
		//LevelDB:    db.DefaultLevelDBConfig(),
		//Listener:   listener.DefaultConfig(),
		//Processors: processor.DefaultConfig(),
	}
}

func (cfg Config) IsValid() (bool, error) {
	if _, err := cfg.Logger.IsValid(); err != nil {
		return false, fmt.Errorf("invalid LoggerConfig: %v", err)
	}
	//if _, err := cfg.Webhooks.IsValid(); err != nil {
	//	return false, fmt.Errorf("invalid WebHookConfig: %v", err)
	//}
	if _, err := cfg.Clients.IsValid(); err != nil {
		return false, fmt.Errorf("invalid service config: %v", err)
	}
	//if _, err := cfg.LevelDB.IsValid(); err != nil {
	//	return false, fmt.Errorf("invalid leveldb config: %v", err)
	//}
	//if _, err := cfg.Listener.IsValid(); err != nil {
	//	return false, fmt.Errorf("invalid listener config: %v", err)
	//}
	//if _, err := cfg.Processors.IsValid(); err != nil {
	//	return false, fmt.Errorf("invalid processors config: %v", err)
	//}

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
