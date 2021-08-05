// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"mall/app/product/service/internal/biz"
	"mall/app/product/service/internal/conf"
	"mall/app/product/service/internal/data"
	"mall/app/product/service/internal/server"
	"mall/app/product/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newApp))
}
