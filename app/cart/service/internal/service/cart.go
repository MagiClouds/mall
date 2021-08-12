package service

import (
    "context"

    pb "mall/api/cart/sevrice/v1"
)

type CartService struct {
	pb.UnimplementedCartServer
}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) AddCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CreateCartReply, error) {
	return &pb.CreateCartReply{}, nil
}
func (s *CartService) IncrItem(ctx context.Context, req *pb.IncrItemRequest) (*pb.IncrItemReply, error) {
	return &pb.IncrItemReply{}, nil
}
func (s *CartService) DecrItem(ctx context.Context, req *pb.DecrItemRequest) (*pb.DecrItemReply, error) {
	return &pb.DecrItemReply{}, nil
}
func (s *CartService) DeleteItemById(ctx context.Context, req *pb.DeleteCartItemRequest) (*pb.DeleteCartItemReply, error) {
	return &pb.DeleteCartItemReply{}, nil
}
func (s *CartService) CleanCart(ctx context.Context, req *pb.CleanCartRequest) (*pb.CleanCartReply, error) {
	return &pb.CleanCartReply{}, nil
}
func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartReply, error) {
	return &pb.GetCartReply{}, nil
}
