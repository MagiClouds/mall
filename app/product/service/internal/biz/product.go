package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ProductDo struct {
	Id          int64
	Name        string
	Sku         string  //单独建表，关联价格相关参数
	Price       int64
	Description string
}

type ProductRepo interface {
	CreateProduct(context.Context, *ProductDo) (int64, error)
	UpdateProduct(context.Context, *ProductDo) (int64, error)
	DeleteProduct(context.Context, int64) error
	GetProduct(context.Context, int64) (*ProductDo, error)
	ListProduct(context.Context) ([]*ProductDo, error)
}

type ProductImage struct {
	Id        int64
	ProductId int64
	Name      string
	Code      string
	Url       string
}

type ImageRepo interface {
	CreateImage(context.Context, []*ProductImage) error
	DeleteImage(context.Context, int64) error
	GetImage(context.Context, []int64) (map[int64][]*ProductImage, error)
}

type ProductSeo struct {
	Id          int64
	ProductId   int64
	Title       string
	Description string
}

type SeoRepo interface {
	CreateSeo(context.Context, []*ProductSeo) error
	UpdateSeo(context.Context, *ProductSeo) error
	DeleteSeo(context.Context, int64) error
	GetSeo(context.Context, []int64) (map[int64]*ProductSeo, error)
}

type ProductSize struct {
	Id        int64
	ProductId int64
	Name      string
	Code      string
}

type SizeRepo interface {
	CreateSize(context.Context, *ProductSize) error
	DeleteSize(context.Context, *ProductSize) error
	GetSize(context.Context, int64) (map[int64][]*ProductSize, error)
}

type ProductBo struct {
	Id          int64
	Name        string
	Sku         string
	Price       int64
	Description string
	Image       []ProductImage
	Size        []ProductSize
	Seo         ProductSeo
}

type ProductUsecase struct {
	productRepo ProductRepo
	imageRepo   ImageRepo
	sizeRepo    SizeRepo
	seoRepo     SeoRepo

	log *log.Helper
}

func NewProductUsecase(productRepo ProductRepo, imageRepo ImageRepo, sizeRepo SizeRepo, seoRepo SeoRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
		imageRepo:   imageRepo,
		sizeRepo:    sizeRepo,
		seoRepo:     seoRepo,
		log:         log.NewHelper(logger),
	}
}

func (uc *ProductUsecase) Create(ctx context.Context, g *ProductBo) (int64, error) {
	//todo 事物
	return uc.productRepo.CreateProduct(ctx, &ProductDo{
		Name:        g.Name,
		Sku:         g.Sku,
		Price:       g.Price,
		Description: g.Description,
	})
}

func (uc *ProductUsecase) Update(ctx context.Context, g *ProductBo) (int64, error) {
	return uc.productRepo.UpdateProduct(ctx, &ProductDo{
		Id:          g.Id,
		Name:        g.Name,
		Sku:         g.Sku,
		Price:       g.Price,
		Description: g.Description,
	})
}

func (uc *ProductUsecase) Delete(ctx context.Context, id int64) error {
	return uc.productRepo.DeleteProduct(ctx, id)
}

func (uc *ProductUsecase) Get(ctx context.Context, id int64) (*ProductBo, error) {
	product, err := uc.productRepo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return &ProductBo{
		Id:          product.Id,
		Name:        product.Name,
		Sku:         "",
		Price:       0,
		Description: "",
		Image:       nil,
		Size:        nil,
		Seo:         ProductSeo{},
	}, nil
}

func (uc *ProductUsecase) List(ctx context.Context) ([]*ProductBo, error) {
	product, err := uc.productRepo.ListProduct(ctx)
	if err != nil {
		return nil, err
	}

	bo := make([]*ProductBo, 0, len(product))
	for _, p := range product {
		bo = append(bo, &ProductBo{
			Id:          p.Id,
			Name:        p.Name,
			Sku:         "",
			Price:       0,
			Description: "",
			Image:       nil,
			Size:        nil,
			Seo:         ProductSeo{},
		})
	}
}
