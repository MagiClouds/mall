package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"mall/app/product/service/internal/conf"
)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(NewData, NewProductRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
