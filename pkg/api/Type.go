package api

import (
	"github.com/simplecontainer/traefik-provider/pkg/configuration"
	"github.com/simplecontainer/traefik-provider/pkg/provider"
)

type Api struct {
	Config   *configuration.Configuration
	Provider *provider.Provider
}
