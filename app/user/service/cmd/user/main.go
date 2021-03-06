package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/registry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	logger2 "mall/pkg/logger"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"mall/app/user/service/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "mall.user.service"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()
	zapLogger := logger2.NewZapLogger(zap.NewProductionEncoderConfig(), zap.NewAtomicLevel())
	zapLogger.Logger = zapLogger.Logger.With(
		zapcore.Field{
			Key:       id,
			Type:      zapcore.NamespaceType,
			Integer:   1,
			String:    Name,
			Interface: Name,
		},
		zapcore.Field{
			Key:       Name,
			Type:      zapcore.NamespaceType,
			Integer:   1,
			String:    Name,
			Interface: Name,
		},
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Server, bc.Data, zapLogger, bc.Registry)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
