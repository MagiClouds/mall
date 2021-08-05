package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Product struct {
	Hello string
}

type ProductRepo interface {
	CreateProduct(context.Context, *Product) error
	UpdateProduct(context.Context, *Product) error
}

type ProductUsecase struct {
	repo ProductRepo
	log  *log.Helper
}

func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ProductUsecase) Create(ctx context.Context, g *Product) error {
	return uc.repo.CreateProduct(ctx, g)
}

func (uc *ProductUsecase) Update(ctx context.Context, g *Product) error {
	return uc.repo.UpdateProduct(ctx, g)
}
