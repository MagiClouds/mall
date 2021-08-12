package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type OrderBo struct {
	Id          int64
	PayStatus   int64
	ShipStatus  int64
	Price       int64
	OrderDetail []*OrderDetail
}

type OrderDetail struct {
	Id            int64
	ProductId     int64
	ProductNum    int64
	ProductSizeId int64
	ProductPrice  int64
	OrderId       int64
}

type OrderRepo interface {
	CreateOrder(context.Context, *OrderBo) error
	UpdateOrder(context.Context, *OrderBo) error
	GetOrder(context.Context, int64) (*OrderBo, error)
	DeleteOrder(context.Context, int64) error
	ListOrder(context.Context) ([]*OrderBo, error)
	UpdateOrderPayStatus(context.Context, *OrderBo) error
	UpdateOrderShipStatus(context.Context, *OrderBo) error
}

type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OrderUsecase) CreateOrder(ctx context.Context, g *OrderBo) error {
	return uc.repo.CreateOrder(ctx, g)
}

func (uc *OrderUsecase) UpdateOrder(ctx context.Context, g *OrderBo) error {
	return uc.repo.UpdateOrder(ctx, g)
}

func (uc *OrderUsecase) GetOrder(ctx context.Context, id int64) (*OrderBo, error) {
	return uc.repo.GetOrder(ctx, id)
}

func (uc *OrderUsecase) DeleteOrder(ctx context.Context, id int64) error {
	return uc.repo.DeleteOrder(ctx, id)
}

func (uc *OrderUsecase) ListOrder(ctx context.Context) ([]*OrderBo, error) {
	return uc.repo.ListOrder(ctx)
}

func (uc *OrderUsecase) UpdateOrderPayStatus(ctx context.Context, g *OrderBo) error {
	return uc.repo.UpdateOrder(ctx, g)
}

func (uc *OrderUsecase) UpdateOrderShipStatus(ctx context.Context, g *OrderBo) error {
	return uc.repo.UpdateOrder(ctx, g)
}
