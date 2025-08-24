package provider

import (
	"context"
	"dario.cat/mergo"
	"github.com/mitchellh/mapstructure"
	"github.com/simplecontainer/smr/pkg/logger"
	"github.com/simplecontainer/traefik-provider/pkg/kinds"
	"github.com/simplecontainer/traefik-provider/pkg/traefik"
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

func (p *Provider) Reload() error {
	return nil
}

func (p *Provider) Mapper(dynamicConfigurations []*kinds.Traefik) {
	tmp := make(map[string]interface{})

	for _, dynamicConfiguration := range dynamicConfigurations {
		err := mergo.Merge(&tmp, dynamicConfiguration.Traefik)

		if err != nil {
			logger.Log.Error("failed to merge traefik configurations", zap.Error(err))
		}
	}

	err := mapstructure.Decode(tmp, p.Config)
	if err != nil {
		logger.Log.Error("failed to decode traefik configurations", zap.Error(err))
	}
}
