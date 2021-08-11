package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/product/service/internal/biz"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p productRepo) CreateProduct(context.Context, *biz.ProductDo) (int64, error) {
	panic("implement me")
}

func (p productRepo) UpdateProduct(context.Context, *biz.ProductDo) (int64, error) {
	panic("implement me")
}
func (p productRepo) DeleteProduct(context.Context, int64) error {
	panic("implement me")
}

func (p productRepo) GetProduct(context.Context, int64) (*biz.ProductDo, error) {
	panic("implement me")
}

func (p productRepo) ListProduct(context.Context) ([]*biz.ProductDo, error) {
	panic("implement me")
}

