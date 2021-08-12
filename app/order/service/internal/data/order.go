package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mall/app/order/service/internal/biz"
)

const (
	tableOrder       = "order"
	tableOrderDetail = "order_detail"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, g *biz.OrderBo) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		id, err := InsertOrder(ctx, tx, &OrderDo{
			PayStatus:  g.PayStatus,
			ShipStatus: g.ShipStatus,
			Price:      g.Price,
		})
		if err != nil {
			return err
		}
		details := make([]*OrderDetail, 0, len(g.OrderDetail))
		for _, d := range g.OrderDetail {
			details = append(details, &OrderDetail{
				ProductId:     d.ProductId,
				ProductNum:    d.ProductNum,
				ProductSizeId: d.ProductSizeId,
				ProductPrice:  d.ProductPrice,
				OrderId:       id,
			})
		}

		if err := InsertOrderDetail(ctx, tx, details); err != nil {
			return err
		}
		return nil
	})
}

type OrderDo struct {
	Id         int64 `gorm:"column:id"`
	PayStatus  int64 `gorm:"column:pay_status"`
	ShipStatus int64 `gorm:"column:ship_status"`
	Price      int64 `gorm:"column:price"`
}

func InsertOrder(ctx context.Context, db *gorm.DB, data *OrderDo) (int64, error) {
	err := db.WithContext(ctx).Table(tableOrder).Create(data).Error
	return data.Id, err
}

type OrderDetail struct {
	Id            int64 `gorm:"column:id"`
	ProductId     int64 `gorm:"column:product_id"`
	ProductNum    int64 `gorm:"column:product_num"`
	ProductSizeId int64 `gorm:"column:product_size_id"`
	ProductPrice  int64 `gorm:"column:product_price"`
	OrderId       int64 `gorm:"column:order_id"`
}

func InsertOrderDetail(ctx context.Context, db *gorm.DB, data []*OrderDetail) error {
	return db.WithContext(ctx).Table(tableOrderDetail).Create(&data).Error
}

func (r *orderRepo) UpdateOrder(ctx context.Context, g *biz.OrderBo) error {
	return nil
}

func (r *orderRepo) GetOrder(ctx context.Context, id int64) (*biz.OrderBo, error) {
	order := new(OrderDo)
	err := r.data.db.WithContext(ctx).Table(tableOrder).Where("id=?", id).First(order).Error
	if err != nil {
		return nil, err
	}

	details := make([]OrderDetail, 0)
	err = r.data.db.WithContext(ctx).Table(tableOrderDetail).Where("order_id=?", id).Find(&details).Error
	if err != nil {
		return nil, err
	}

	detailBo := make([]*biz.OrderDetail, 0, len(details))
	for _, d := range details {
		detailBo = append(detailBo, &biz.OrderDetail{
			Id:            d.Id,
			ProductId:     d.ProductId,
			ProductNum:    d.ProductNum,
			ProductSizeId: d.ProductSizeId,
			ProductPrice:  d.ProductPrice,
			OrderId:       d.OrderId,
		})
	}

	return &biz.OrderBo{
		Id:          order.Id,
		PayStatus:   order.PayStatus,
		ShipStatus:  order.ShipStatus,
		Price:       order.Price,
		OrderDetail: detailBo,
	}, nil
}

func (r *orderRepo) DeleteOrder(ctx context.Context, id int64) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Table(tableOrder).Where("id=?", id).Delete(&OrderDo{}).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Table(tableOrderDetail).Where("order_id=?", id).Delete(&OrderDetail{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *orderRepo) ListOrder(context.Context) ([]*biz.OrderBo, error) {
	panic("implement me")
}

func (r *orderRepo) UpdateOrderPayStatus(context.Context, *biz.OrderBo) error {
	panic("implement me")
}

func (r *orderRepo) UpdateOrderShipStatus(context.Context, *biz.OrderBo) error {
	panic("implement me")
}
