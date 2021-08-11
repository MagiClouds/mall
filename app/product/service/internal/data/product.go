package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mall/app/product/service/internal/biz"
)

const tableProduct = "product"

type productRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p productRepo) CreateProduct(ctx context.Context, product *biz.ProductDo) (int64, error) {
	res := p.data.db.WithContext(ctx).Table(tableProduct).Create(product)
	return product.Id, res.Error
}

func (p productRepo) UpdateProduct(ctx context.Context, product *biz.ProductDo) (int64, error) {
	res := p.data.db.WithContext(ctx).Table(tableProduct).Where("id=?", product.Id).Updates(product)
	return product.Id, res.Error
}
func (p productRepo) DeleteProduct(ctx context.Context, id int64) error {
	res := p.data.db.WithContext(ctx).Table(tableProduct).Where("id=?", id).Delete(biz.ProductDo{})
	return res.Error
}

func (p productRepo) GetProduct(ctx context.Context, id int64) (*biz.ProductDo, error) {
	product := new(biz.ProductDo)
	res := p.data.db.WithContext(ctx).Table(tableProduct).Where("id=?", id).First(product)
	return product, res.Error
}

func (p productRepo) ListProduct(context.Context) ([]*biz.ProductDo, error) {
	panic("implement me")
}
