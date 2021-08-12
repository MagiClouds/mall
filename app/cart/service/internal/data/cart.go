package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mall/app/cart/service/internal/biz"
)

const tableCart = "cart"

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *cartRepo) AddCart(ctx context.Context, data *biz.Cart) error {
	return r.data.db.WithContext(ctx).Table(tableCart).Create(data).Error
}

func (r *cartRepo) IncrItem(ctx context.Context, i *biz.Item) error {
	return r.data.db.WithContext(ctx).Table(tableCart).Where("user_id=? and product_id=?", i.UserId, i.ItemId).
		Update("num", gorm.Expr("num+?", i.ChangeNum)).Error
}

func (r *cartRepo) DecrItem(ctx context.Context, i *biz.Item) error {
	//todo num > 0
	return r.data.db.WithContext(ctx).Table(tableCart).Where("user_id=? and product_id=?", i.UserId, i.ItemId).
		Update("num", gorm.Expr("num-?", i.ChangeNum)).Error
}

func (r *cartRepo) DeleteItemById(ctx context.Context, i *biz.Item) error {
	return r.data.db.WithContext(ctx).Table(tableCart).Where("user_id=? and product_id=?", i.UserId, i.ItemId).
		Delete(&biz.Cart{}).Error
}

func (r *cartRepo) CleanCart(ctx context.Context, id int64) error {
	return r.data.db.WithContext(ctx).Table(tableCart).Where("user_id=?", id).Delete(&biz.Cart{}).Error
}

func (r *cartRepo) GetCart(ctx context.Context, id int64) ([]*biz.Cart, error) {
	bo := make([]*biz.Cart, 0)
	return bo, r.data.db.WithContext(ctx).Table(tableCart).Where("user_id=?", id).Find(&bo).Error
}
