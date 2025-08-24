package traefik

import (
	"context"
	"fmt"
	"github.com/simplecontainer/smr/pkg/configuration"
	"github.com/simplecontainer/smr/pkg/logger"
	"github.com/simplecontainer/traefik-provider/pkg/kinds"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type EtcdProvider struct {
	client *clientv3.Client
	logger *zap.Logger
	prefix string
}

func NewEtcdProvider(endpoint string, prefix string) (*EtcdProvider, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: configuration.Timeout.EtcdConnectionTimeout,
	})

	if err != nil {
		return nil, err
	}

	return &EtcdProvider{
		client: client,
		logger: logger.Log,
		prefix: prefix,
	}, nil
}

func (e *EtcdProvider) Name() string {
	return "etcd"
}

func (e *EtcdProvider) GetServices(ctx context.Context) ([]*kinds.Traefik, error) {
	resp, err := e.client.Get(ctx, e.prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to get services from etcd: %w", err)
	}

	var configurations []*kinds.Traefik = nil

	for _, kv := range resp.Kvs {
		config, err := kinds.New(kv.Value)
		if err != nil {
			logger.Log.Error("failed to parse traefik custom resource", zap.Error(err))
			continue
		}

		configurations = append(configurations, config)
	}

	return configurations, nil
}

func (e *EtcdProvider) Watch(ctx context.Context, callback func([]*kinds.Traefik)) error {
	watchChan := e.client.Watch(ctx, e.prefix, clientv3.WithPrefix())

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			e.logger.Debug("etcd watch event", zap.String("type", event.Type.String()), zap.String("key", string(event.Kv.Key)))
		}

		services, err := e.GetServices(ctx)
		if err != nil {
			e.logger.Error("failed to get services during watch", zap.Error(err))
			continue
		}

		callback(services)
	}

	return nil
}
