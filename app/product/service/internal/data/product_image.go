package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/product/service/internal/biz"
)

type imageRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewImageRepo(data *Data, logger log.Logger) biz.ImageRepo {
	return &imageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (i imageRepo) CreateImage(context.Context, []*biz.ProductImage) error {
	panic("implement me")
}

func (i imageRepo) DeleteImage(context.Context, int64) error {
	panic("implement me")
}

func (i imageRepo) GetImage(context.Context, []int64) (map[int64][]*biz.ProductImage, error) {
	panic("implement me")
}


