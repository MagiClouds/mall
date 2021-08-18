package server

import (
	"github.com/golang/protobuf/ptypes/duration"
	"mall/app/user/service/internal/conf"
	"testing"
)

func TestNewRegistrar(t *testing.T) {
	NewRegistrar(&conf.Registry{Etcd: &conf.Registry_Etcd{
		Endpoint: []string{"127.0.0.1:2379"},
		Timeout:  &duration.Duration{Seconds: 10},
	}})
}
