package server

import (
	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"mall/app/user/service/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) (registry.Registrar, func(), error) {
	client, err := clientv3.New(clientv3.Config{Endpoints: conf.Etcd.Endpoint,
		DialTimeout: conf.Etcd.Timeout.AsDuration(), DialOptions: []grpc.DialOption{grpc.WithBlock()}})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		client.Close()
	}

	r := etcd.New(client)
	return r, func() {
		cleanup()
	}, nil
}
