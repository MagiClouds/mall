package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id    int
	Name  string
	Phone string
	Pwd   string
}

type UserLoginBo struct {
	Id    int
	Name  string
	Token string
}

type ListUserDto struct{
	LastTime int64
	Rn       int64
}

type UserRepo interface {
	UserRegister(context.Context, *User) error
	UserLogin(context.Context, *User) (*UserLoginBo, error)
	UpdateUser(context.Context, *User) error
	DeleteUser(context.Context, int64) error
	GetUser(context.Context, int64) (*User, error)
	ListUser(context.Context, *ListUserDto) ([]*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, g *User) error {
	return uc.repo.UserRegister(ctx, g)
}

func (uc *UserUsecase) Update(ctx context.Context, g *User) error {
	return uc.repo.UpdateUser(ctx, g)
}
