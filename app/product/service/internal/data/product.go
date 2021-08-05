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

func (r *productRepo) CreateProduct(ctx context.Context, g *biz.Product) error {
	return nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, g *biz.Product) error {
	return nil
}
