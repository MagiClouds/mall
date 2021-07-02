package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mall/app/user/service/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

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

const tableUser = "user"

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) Save(ctx context.Context, user *biz.UserDto) error {
	res := u.data.db.WithContext(ctx).Table(tableUser).Create(&UserDo{Name: user.Name,
		Phone: user.Phone, Password: user.Pwd})
	return res.Error
}

func (u userRepo) GetByPhone(ctx context.Context, phone string) (*biz.UserDo, error) {
	user := new(UserDo)
	res := u.data.db.WithContext(ctx).Table(tableUser).First(user, "phone = ?", phone)

	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return new(biz.UserDo), nil
		}
		return nil, res.Error
	}

	return translateDataBoToBizBo(user), nil
}

func (u userRepo) Update(ctx context.Context, user *biz.UserDto) error {
	res := u.data.db.WithContext(ctx).Where("id = ?", user.Id).Table(tableUser).Updates(UserDo{
		Name: user.Name,
	})

	return res.Error
}

func (u userRepo) Delete(ctx context.Context, id int64) error {
	res := u.data.db.WithContext(ctx).Table(tableUser).Delete(&UserDo{Id: id})
	return res.Error
}

func (u userRepo) GetById(ctx context.Context, id int64) (*biz.UserDo, error) {
	user := new(UserDo)
	res := u.data.db.WithContext(ctx).Table(tableUser).Where("id = ?", id).First(user)

	if res.Error == gorm.ErrRecordNotFound {
		return new(biz.UserDo), nil
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return translateDataBoToBizBo(user), nil
}

func (u userRepo) List(ctx context.Context, filter *biz.ListUserDto) ([]*biz.UserDo, error) {
	if filter.Rn == 0 || filter.Rn > 20 {
		filter.Rn = 10
	}

	bo := make([]*biz.UserDo, 0, filter.Rn)

	users := make([]*UserDo, 0, filter.Rn)
	res := u.data.db.WithContext(ctx).Table(tableUser).
		Where("create_time < ?", filter.LastTime).
		Order("create_time desc").Limit(int(filter.Rn)).Find(&users)
	if res.Error != nil {
		return bo, res.Error
	}

	for _, u := range users {
		bo = append(bo, translateDataBoToBizBo(u))
	}

	return bo, nil
}

func translateDataBoToBizBo(do *UserDo) *biz.UserDo {
	return &biz.UserDo{
		Id:    do.Id,
		Name:  do.Name,
		Phone: do.Phone,
		Pwd:   do.Password,
	}
}
