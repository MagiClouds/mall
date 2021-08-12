package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/cart/service/internal/biz"
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

func (r *cartRepo) AddCart(context.Context, *biz.Cart) error {
	panic("implement me")
}

func (r *cartRepo) IncrItem(context.Context, *biz.Item) error {
	panic("implement me")
}

func (r *cartRepo) DecrItem(context.Context, *biz.Item) error {
	panic("implement me")
}

func (r *cartRepo) DeleteItemById(context.Context, *biz.Item) error {
	panic("implement me")
}

func (r *cartRepo) CleanCart(context.Context, int64) error {
	panic("implement me")
}

func (r *cartRepo) GetCart(context.Context, int64) []*biz.Cart {
	panic("implement me")
}
