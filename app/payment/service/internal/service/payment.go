package service

import (
	"context"

	v1 "mall/api/payment/service/v1"
	"mall/app/payment/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedPaymentServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.CreatePaymentRequest) (*v1.CreatePaymentReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.String())

	if in.String() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.String())
	}
	return &v1.CreatePaymentReply{}, nil
}
