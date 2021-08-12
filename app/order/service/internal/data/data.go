package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall/app/order/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.Source, // DSN data source name
		DefaultStringSize:         256,               // string 类型字段的默认长度
		DisableDatetimePrecision:  true,              // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,              // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,              // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	cleanup := func() {
		helper.Info("closing the data resources")
	}

	return &Data{db: db}, cleanup, nil
}
