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
