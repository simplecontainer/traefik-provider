package provider

import (
	"context"
	"dario.cat/mergo"
	"encoding/json"
	"github.com/simplecontainer/smr/pkg/logger"
	"github.com/simplecontainer/traefik-provider/pkg/kinds"
	"github.com/simplecontainer/traefik-provider/pkg/traefik"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"go.uber.org/zap"
)

func New(endpoint string) (*Provider, error) {
	provider, err := traefik.NewEtcdProvider(endpoint, "/traefik.io/v1/kind/custom/")
	if err != nil {
		return nil, err
	}

	return &Provider{
		Provider: provider,
		Config:   traefik.NewConfigurationManager(),
	}, nil
}

func (p *Provider) Watch() error {
	return p.Provider.Watch(context.Background(), p.Mapper)
}

func (p *Provider) Mapper(dynamicConfigurations []*kinds.Traefik) {
	if dynamicConfigurations == nil {
		p.Config = &dynamic.Configuration{}
	} else {
		tmp := make(map[string]interface{})

		for _, dynamicConfiguration := range dynamicConfigurations {
			err := mergo.Merge(&tmp, dynamicConfiguration.Traefik)

			if err != nil {
				logger.Log.Error("failed to merge traefik dynamic configurations", zap.Error(err))
			}
		}

		bytes, err := json.Marshal(tmp)
		if err != nil {
			logger.Log.Error("failed to marshal traefik dynamic configuration", zap.Error(err))
		}

		err = json.Unmarshal(bytes, p.Config)
		if err != nil {
			logger.Log.Error("failed to decode traefik dynamic configuration", zap.Error(err))
		}
	}
}
