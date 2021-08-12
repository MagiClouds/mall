package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/cart/service/internal/biz"

	pb "mall/api/cart/sevrice/v1"
)

type CartService struct {
	pb.UnimplementedCartServer

	uc  *biz.CartUsecase
	log *log.Helper
}

func NewCartService(uc *biz.CartUsecase, logger log.Logger) *CartService {
	return &CartService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *CartService) AddCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CreateCartReply, error) {
	err := s.uc.AddCart(ctx, &biz.Cart{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		SizeId:    req.SizeId,
		Num:       req.Num,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateCartReply{Message: "success"}, nil
}
func (s *CartService) IncrItem(ctx context.Context, req *pb.IncrItemRequest) (*pb.IncrItemReply, error) {
	err := s.uc.IncrItem(ctx, &biz.Item{
		UserId:    req.UserId,
		ItemId:    req.ItemId,
		ChangeNum: req.ChangeNum,
	})
	return &pb.IncrItemReply{Message: "success"}, err
}
func (s *CartService) DecrItem(ctx context.Context, req *pb.DecrItemRequest) (*pb.DecrItemReply, error) {
	err := s.uc.DecrItem(ctx, &biz.Item{
		UserId:    req.UserId,
		ItemId:    req.ItemId,
		ChangeNum: req.ChangeNum,
	})
	return &pb.DecrItemReply{Message: "success"}, err
}
func (s *CartService) DeleteItemById(ctx context.Context, req *pb.DeleteCartItemRequest) (*pb.DeleteCartItemReply, error) {
	err := s.uc.DeleteItemById(ctx, &biz.Item{
		UserId: req.UserId,
		ItemId: req.ItemId,
	})
	return &pb.DeleteCartItemReply{Message: "success"}, err
}
func (s *CartService) CleanCart(ctx context.Context, req *pb.CleanCartRequest) (*pb.CleanCartReply, error) {
	err := s.uc.CleanCart(ctx, req.Id)
	return &pb.CleanCartReply{Message: "success"}, err
}
func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartReply, error) {
	data, err := s.uc.GetCart(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	cart := make([]*pb.GetCartReply_Cart, 0, len(data))
	for _, d := range data {
		cart = append(cart, &pb.GetCartReply_Cart{
			Id:        d.Id,
			ProductId: d.ProductId,
			SizeId:    d.SizeId,
			Num:       d.Num,
		})
	}

	return &pb.GetCartReply{
		Cart: cart,
	}, nil
}
