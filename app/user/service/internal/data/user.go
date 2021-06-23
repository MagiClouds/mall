package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/user/service/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) UserRegister(context.Context, *biz.User) error {
	panic("implement me")
}

func (u userRepo) UserLogin(context.Context, *biz.User) (*biz.UserLoginBo, error) {
	panic("implement me")
}

func (u userRepo) UpdateUser(context.Context, *biz.User) error {
	panic("implement me")
}

func (u userRepo) DeleteUser(context.Context, int64) error {
	panic("implement me")
}

func (u userRepo) GetUser(context.Context, int64) (*biz.User, error) {
	panic("implement me")
}

func (u userRepo) ListUser(context.Context, *biz.ListUserDto) ([]*biz.User, error) {
	panic("implement me")
}


