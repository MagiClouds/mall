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
	order, err := s.uc.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	orderDetailVo := make([]*pb.OrderDetail, 0, len(order.OrderDetail))
	for _, d := range order.OrderDetail {
		orderDetailVo = append(orderDetailVo, &pb.OrderDetail{
			Id:            d.Id,
			ProductId:     d.ProductId,
			ProductNum:    d.ProductNum,
			ProductSizeId: d.ProductSizeId,
			ProductPrice:  d.ProductPrice,
			OrderId:       d.OrderId,
		})
	}

	return &pb.GetOrderReply{
		Id:          order.Id,
		PayStatus:   order.PayStatus,
		ShipStatus:  order.ShipStatus,
		Price:       order.Price,
		OrderDetail: orderDetailVo,
	}, nil
}
func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	return &pb.ListOrderReply{}, nil
}
func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	orderDetailBo := make([]*biz.OrderDetail, 0, len(req.OrderDetail))
	for _, d := range req.OrderDetail {
		orderDetailBo = append(orderDetailBo, &biz.OrderDetail{
			ProductId:     d.ProductId,
			ProductNum:    d.ProductNum,
			ProductSizeId: d.ProductSizeId,
			ProductPrice:  d.ProductPrice,
		})
	}

	if err := s.uc.CreateOrder(ctx, &biz.OrderBo{
		PayStatus:   req.PayStatus,
		ShipStatus:  req.ShipStatus,
		Price:       req.Price,
		OrderDetail: orderDetailBo,
	}); err != nil {
		return nil, err
	}

	return &pb.CreateOrderReply{Message: "success"}, nil
}
func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderReply, error) {
	err := s.uc.DeleteOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteOrderReply{Message: "success"}, nil
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
