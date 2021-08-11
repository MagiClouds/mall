package service

import (
    "context"
	"mall/app/order/service/internal/biz"

	pb "mall/api/order/service/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer

	uc *biz.OrderUsecase
}

func NewOrderService(uc *biz.OrderUsecase) *OrderService {
	return &OrderService{uc: uc}
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	return &pb.GetOrderReply{}, nil
}
func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	return &pb.ListOrderReply{}, nil
}
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	return &pb.CreateOrderReply{}, nil
}
func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderReply, error) {
	return &pb.DeleteOrderReply{}, nil
}
func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderReply, error) {
	return &pb.UpdateOrderReply{}, nil
}
func (s *OrderService) UpdateOrderPayStatus(ctx context.Context, req *pb.UpdateOrderPayStatusRequest) (*pb.UpdateOrderPayStatusReply, error) {
	return &pb.UpdateOrderPayStatusReply{}, nil
}
func (s *OrderService) UpdateOrderShipStatus(ctx context.Context, req *pb.UpdateOrderShipStatusRequest) (*pb.UpdateOrderShipStatusReply, error) {
	return &pb.UpdateOrderShipStatusReply{}, nil
}
