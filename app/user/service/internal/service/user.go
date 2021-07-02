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
	if err := s.uc.UserRegister(ctx, &biz.UserDto{
		Name:  req.UserName,
		Phone: req.Phone,
		Pwd:   req.Pwd,
	}); err != nil {
		return nil, err
	}

	return &pb.BaseResponse{Message: "success"}, nil
}
func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	user, err := s.uc.UserLogin(ctx, &biz.UserDto{
		Phone: req.GetPhone(),
		Pwd:   req.GetPwd(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UserLoginResponse{
		Id:    user.Id,
		Name:  user.Name,
		Token: user.Token,
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.BaseResponse, error) {
	err := s.uc.UpdateUser(ctx, &biz.UserDto{
		Id:    req.Id,
		Name:  req.UserName,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &pb.BaseResponse{Message: "success"}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.BaseResponse, error) {
	if err := s.uc.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &pb.BaseResponse{Message: "success"}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	user, err := s.uc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserInfo{
		Id:       user.Id,
		UserName: user.Name,
		Phone:    user.Phone,
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	user, err := s.uc.ListUser(ctx, &biz.ListUserDto{
		LastTime: req.LastTime,
		Rn:       req.Rn,
	})
	if err != nil {
		return nil, err
	}

	vo := make([]*pb.UserInfo, 0, len(user))
	for _, u := range user {
		vo = append(vo, &pb.UserInfo{
			Id:       u.Id,
			UserName: u.Name,
			Phone:    u.Phone,
		})
	}

	return &pb.ListUserResponse{User: vo}, nil
}
