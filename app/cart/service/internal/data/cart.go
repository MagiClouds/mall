package data

import (
	"mall/app/cart/service/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *cartRepo) CreateCart(ctx context.Context, g *biz.Cart) error {
	return nil
}

func (r *cartRepo) UpdateCart(ctx context.Context, g *biz.Cart) error {
	return nil
}
