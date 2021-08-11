package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/product/service/internal/biz"

	pb "mall/api/product/service/v1"
)

type ProductService struct {
	pb.UnimplementedProductServer

	puc *biz.ProductUsecase
	log *log.Helper
}

func NewProductService(uc *biz.ProductUsecase, logger log.Logger) *ProductService {
	return &ProductService{
		puc: uc,

		log: log.NewHelper(logger),
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	image := make([]biz.ProductImage, 0, len(req.Image))
	for _, i := range req.Image {
		image = append(image, biz.ProductImage{
			Name: i.Name,
			Url:  i.Url,
			Code: i.Code,
		})
	}

	size := make([]biz.ProductSize, 0, len(req.Size))
	for _, s := range req.Size {
		size = append(size, biz.ProductSize{
			Name: s.Name,
			Code: s.Code,
		})
	}

	id, err := s.puc.Create(ctx, &biz.ProductBo{
		Name:        req.Name,
		Sku:         req.Sku,
		Price:       req.Price,
		Description: req.Description,
		Image:       image,
		Size:        size,
		Seo: biz.ProductSeo{
			Title:       req.Seo.Title,
			Description: req.Seo.Description,
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductReply{Id: id}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	image := make([]biz.ProductImage, 0, len(req.Image))
	for _, i := range req.Image {
		image = append(image, biz.ProductImage{
			Id:        i.Id,
			ProductId: req.Id,
			Name:      i.Name,
			Code:      i.Code,
			Url:       i.Url,
		})
	}

	size := make([]biz.ProductSize, 0, len(req.Size))
	for _, s := range req.Size {
		size = append(size, biz.ProductSize{
			Id:        s.Id,
			ProductId: req.Id,
			Name:      s.Name,
			Code:      s.Code,
		})
	}

	id, err := s.puc.Update(ctx, &biz.ProductBo{
		Id:          req.Id,
		Name:        req.Name,
		Sku:         req.Sku,
		Price:       req.Price,
		Description: req.Description,
		Image:       image,
		Size:        size,
		Seo: biz.ProductSeo{
			Title:       req.Seo.Title,
			Description: req.Seo.Description,
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProductReply{Id: id}, nil
}
func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	err := s.puc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProductReply{Message: "success"}, nil
}
func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	product, err := s.puc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductReply{
		Id:          product.Id,
		Name:        product.Name,
		Sku:         "",
		Price:       0,
		Description: "",
		Image:       nil,
		Size:        nil,
		Seo:         nil,
	}, nil
}
func (s *ProductService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductReply, error) {
	product, err := s.puc.List(ctx)
	if err != nil {
		return nil, err
	}

	vo := make([]*pb.GetProductReply, 0, len(product))

	for _, p := range product {
		vo = append(vo, &pb.GetProductReply{
			Id:          p.Id,
			Name:        p.Name,
			Sku:         "",
			Price:       0,
			Description: "",
			Image:       nil,
			Size:        nil,
			Seo:         nil,
		})
	}

	return &pb.ListProductReply{Product: vo}, nil
}
