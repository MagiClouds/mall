package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/product/service/internal/biz"
)

type seoRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewSeoRepo(data *Data, logger log.Logger) biz.SeoRepo {
	return &seoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s seoRepo) CreateSeo(context.Context, []*biz.ProductSeo) error {
	panic("implement me")
}

func (s seoRepo) UpdateSeo(context.Context, *biz.ProductSeo) error {
	panic("implement me")
}

func (s seoRepo) DeleteSeo(context.Context, int64) error {
	panic("implement me")
}

func (s seoRepo) GetSeo(context.Context, []int64) (map[int64]*biz.ProductSeo, error) {
	panic("implement me")
}


