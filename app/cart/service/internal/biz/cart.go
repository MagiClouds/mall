package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Cart struct {
	Hello string
}

type Item struct {
	userId    int64
	IterId    int64
	ChangeNum int64
}

type CartRepo interface {
	AddCart(context.Context, *Cart) error
	IncrItem(context.Context, *Item) error
	DecrItem(context.Context, *Item) error
	DeleteItemById(context.Context, *Item) error
	CleanCart(context.Context, int64) error
	GetCart(context.Context, int64) []*Cart
}

type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) Create(ctx context.Context, g *Cart) error {
	return uc.repo.AddCart(ctx, g)
}

func (uc *CartUsecase) CleanCart(ctx context.Context, id int64) error {
	return uc.repo.CleanCart(ctx, id)
}
