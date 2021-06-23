package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall/app/user/service/internal/biz"
	"os"
	"testing"
)


func Test_userRepo_UserRegister(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	data := Data{db:db}
	l := log.With(log.NewStdLogger(os.Stdout))
	u := NewUserRepo(&data, l)
	user := biz.User{Name:"1", Phone:"13162568196", Pwd:"123456"}
	if err := u.UserRegister(context.Background(), &user); err != nil {
		t.Error(err)
	}
}