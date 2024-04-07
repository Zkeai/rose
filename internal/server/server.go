package server

import (
	"github.com/Zkeai/MuCoinPay/McPay-go/common/net/chttp"
	"github.com/Zkeai/MuCoinPay/McPay-go/internal/conf"
	"github.com/Zkeai/MuCoinPay/McPay-go/internal/handler"
	"github.com/Zkeai/MuCoinPay/McPay-go/internal/service"
)

func NewHTTP(conf *conf.Conf) *chttp.Server {
	s := chttp.NewServer((*chttp.Config)(conf.Server))
	svc := service.NewService(conf)

	handler.InitRouter(s, svc)

	err := s.Start()
	if err != nil {
		panic(err)
	}

	return s
}
