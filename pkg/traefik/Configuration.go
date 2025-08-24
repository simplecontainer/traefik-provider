package traefik

import (
	"encoding/json"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func NewConfigurationManager() *dynamic.Configuration {
	return &dynamic.Configuration{}
}

func (cm *ConfigurationManager) GetConfiguration() (*dynamic.Configuration, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	configJSON, _ := json.Marshal(cm.config)
	var config *dynamic.Configuration

	err := json.Unmarshal(configJSON, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
