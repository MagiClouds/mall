package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/order/service/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, g *biz.OrderDo) error {
	return nil
}

func (r *orderRepo) UpdateOrder(ctx context.Context, g *biz.OrderDo) error {
	return nil
}

func (r *orderRepo) GetOrder(context.Context, int64) (*biz.OrderDo, error) {
	panic("implement me")
}

func (r *orderRepo) DeleteOrder(context.Context, int64) error {
	panic("implement me")
}

func (r *orderRepo) ListOrder(context.Context) ([]*biz.OrderDo, error) {
	panic("implement me")
}

func (r *orderRepo) UpdateOrderPayStatus(context.Context, *biz.OrderDo) error {
	panic("implement me")
}

func (r *orderRepo) UpdateOrderShipStatus(context.Context, *biz.OrderDo) error {
	panic("implement me")
}
