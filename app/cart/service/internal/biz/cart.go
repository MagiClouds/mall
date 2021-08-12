package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Cart struct {
	Id        int64 `gorm:"column:id"`
	UserId    int64 `gorm:"column:user_id"`
	ProductId int64 `gorm:"column:product_id"`
	SizeId    int64 `gorm:"column:size_id"`
	Num       int64 `gorm:"column:num"`
}

type Item struct {
	UserId    int64
	ItemId    int64
	ChangeNum int64
}

type CartRepo interface {
	AddCart(context.Context, *Cart) error
	IncrItem(context.Context, *Item) error
	DecrItem(context.Context, *Item) error
	DeleteItemById(context.Context, *Item) error
	CleanCart(context.Context, int64) error
	GetCart(context.Context, int64) ([]*Cart, error)
}

type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) AddCart(ctx context.Context, g *Cart) error {
	//todo 已存在则为增加数量
	return uc.repo.AddCart(ctx, g)
}

func (uc *CartUsecase) IncrItem(ctx context.Context, i *Item) error {
	return uc.repo.IncrItem(ctx, i)
}

func (uc *CartUsecase) DecrItem(ctx context.Context, i *Item) error {
	return uc.repo.DecrItem(ctx, i)
}

func (uc *CartUsecase) DeleteItemById(ctx context.Context, i *Item) error {
	return uc.repo.DeleteItemById(ctx, i)
}

func (uc *CartUsecase) CleanCart(ctx context.Context, id int64) error {
	return uc.repo.CleanCart(ctx, id)
}

func (uc *CartUsecase) GetCart(ctx context.Context, id int64) ([]*Cart, error) {
	return uc.repo.GetCart(ctx, id)
}
