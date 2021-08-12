// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mall/app/cart/service/internal/biz"
	"mall/app/cart/service/internal/conf"
	"mall/app/cart/service/internal/data"
	"mall/app/cart/service/internal/server"
	"mall/app/cart/service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/go-kratos/kratos/v2"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
