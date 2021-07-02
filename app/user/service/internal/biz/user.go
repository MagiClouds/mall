package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type UserDto struct {
	Id    int64
	Name  string
	Phone string
	Pwd   string
}

type UserDo struct {
	Id    int64
	Name  string
	Phone string
	Pwd   string
}

type UserLoginBo struct {
	Id    int64
	Name  string
	Token string
}

type UserBo struct {
	Id    int64
	Name  string
	Phone string
}

type ListUserDto struct {
	LastTime uint64
	Rn       uint32
}

type UserRepo interface {
	Save(context.Context, *UserDto) error
	GetByPhone(context.Context, string) (*UserDo, error)
	Update(context.Context, *UserDto) error
	Delete(context.Context, int64) error
	GetById(context.Context, int64) (*UserDo, error)
	List(context.Context, *ListUserDto) ([]*UserDo, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) UserRegister(ctx context.Context, u *UserDto) error {
	user, err := uc.repo.GetByPhone(ctx, u.Phone)
	if err != nil {
		return err
	}

	if user.Id != 0 {
		return fmt.Errorf("user already exists")
	}

	return uc.repo.Save(ctx, u)
}

func (uc *UserUsecase) UserLogin(ctx context.Context, g *UserDto) (*UserLoginBo, error) {
	user, err := uc.repo.GetByPhone(ctx, g.Phone)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not exists")
	}

	if user.Pwd != g.Pwd {
		return nil, fmt.Errorf("password mistake")
	}

	res := &UserLoginBo{
		Id:    user.Id,
		Name:  user.Name,
		Token: "111",
	}

	return res, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, u *UserDto) error {
	user, err := uc.repo.GetById(ctx, u.Id)
	if err != nil {
		return err
	}

	if user.Id == 0 {
		return fmt.Errorf("user not exist")
	}

	return uc.repo.Update(ctx, u)
}

func (uc *UserUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*UserBo, error) {
	user, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not exists")
	}

	res := &UserBo{Id: user.Id, Name: user.Name, Phone: user.Phone}

	return res, nil
}

func (uc *UserUsecase) ListUser(ctx context.Context, dto *ListUserDto) ([]*UserBo, error) {
	user, err := uc.repo.List(ctx, dto)
	if err != nil {
		uc.log.Errorf("uc.repo.List err: %v", err)
		return nil, fmt.Errorf("system error")
	}

	res := make([]*UserBo, 0, len(user))

	for _, u := range user {
		res = append(res, &UserBo{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		})
	}

	return res, nil
}
