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

type UserDo struct {
	Id       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Phone    string `gorm:"column:phone"`
	Password string `gorm:"column:password"`
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) UserRegister(ctx context.Context, user *biz.User) error {
	res := u.data.db.WithContext(ctx).Table("user").Create(&UserDo{Name: user.Name,
		Phone: user.Phone, Password: user.Pwd})
	return res.Error
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
