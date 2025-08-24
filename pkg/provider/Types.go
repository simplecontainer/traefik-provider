package provider

import (
	"github.com/simplecontainer/traefik-provider/pkg/traefik"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

type Provider struct {
	Provider traefik.Provider
	Config   *dynamic.Configuration
}
