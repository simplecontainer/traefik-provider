package traefik

import (
	"encoding/json"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/tls"
)

func NewConfigurationManager() *dynamic.Configuration {
	return &dynamic.Configuration{
		HTTP: &dynamic.HTTPConfiguration{
			Routers:           make(map[string]*dynamic.Router),
			Services:          make(map[string]*dynamic.Service),
			Middlewares:       make(map[string]*dynamic.Middleware),
			ServersTransports: make(map[string]*dynamic.ServersTransport),
		},
		TCP: &dynamic.TCPConfiguration{
			Routers:           make(map[string]*dynamic.TCPRouter),
			Services:          make(map[string]*dynamic.TCPService),
			Middlewares:       make(map[string]*dynamic.TCPMiddleware),
			ServersTransports: make(map[string]*dynamic.TCPServersTransport),
		},
		UDP: &dynamic.UDPConfiguration{
			Routers:  make(map[string]*dynamic.UDPRouter),
			Services: make(map[string]*dynamic.UDPService),
		},
		TLS: &dynamic.TLSConfiguration{
			Certificates: make([]*tls.CertAndStores, 0),
			Options:      make(map[string]tls.Options),
			Stores:       make(map[string]tls.Store),
		},
	}
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
