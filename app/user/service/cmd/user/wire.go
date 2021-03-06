// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"mall/app/user/service/internal/biz"
	"mall/app/user/service/internal/conf"
	"mall/app/user/service/internal/data"
	"mall/app/user/service/internal/server"
	"mall/app/user/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
