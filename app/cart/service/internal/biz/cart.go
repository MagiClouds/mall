package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Cart struct {
	Hello string
}

type CartRepo interface {
	CreateCart(context.Context, *Cart) error
	UpdateCart(context.Context, *Cart) error
}

type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) Create(ctx context.Context, g *Cart) error {
	return uc.repo.CreateCart(ctx, g)
}

func (uc *CartUsecase) Update(ctx context.Context, g *Cart) error {
	return uc.repo.UpdateCart(ctx, g)
}
