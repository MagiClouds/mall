package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/product/service/internal/biz"
)

type sizeRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewSizeRepo(data *Data, logger log.Logger) biz.SizeRepo {
	return &sizeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sizeRepo) CreateSize(context.Context, *biz.ProductSize) error {
	panic("implement me")
}

func (s sizeRepo) DeleteSize(context.Context, *biz.ProductSize) error {
	panic("implement me")
}

func (s sizeRepo) GetSize(context.Context, int64) (map[int64][]*biz.ProductSize, error) {
	panic("implement me")
}
