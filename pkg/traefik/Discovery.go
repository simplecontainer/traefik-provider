package traefik

import (
	"context"
	"github.com/simplecontainer/traefik-provider/pkg/kinds"
)

type Provider interface {
	GetServices(ctx context.Context) ([]*kinds.Traefik, error)
	Watch(ctx context.Context, callback func([]*kinds.Traefik)) error
	Name() string
}

type DiscoveredService struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Address  string            `json:"address"`
	Port     int               `json:"port"`
	Tags     []string          `json:"tags"`
	Labels   map[string]string `json:"labels"`
	Health   string            `json:"health"`
	Protocol string            `json:"protocol"`
	Path     string            `json:"path,omitempty"`
	Weight   int               `json:"weight,omitempty"`

	TraefikLabels map[string]string `json:"traefikLabels"`
}
