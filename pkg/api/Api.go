package api

import (
	"github.com/simplecontainer/traefik-provider/pkg/configuration"
	"github.com/simplecontainer/traefik-provider/pkg/provider"
)

func New(config *configuration.Configuration) *Api {
	providerManager, err := provider.New(config.Endpoint)
	if err != nil {
		panic(err)
	}

	return &Api{
		Provider: providerManager,
		Config:   config,
	}
}
