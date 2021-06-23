package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/user/service/internal/biz"

	pb "mall/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.BaseResponse, error) {
	return &pb.BaseResponse{}, nil
}
func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.BaseResponse, error) {
	return &pb.BaseResponse{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.BaseResponse, error) {
	return &pb.BaseResponse{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}
