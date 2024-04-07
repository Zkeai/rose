package handler

import (
	"github.com/Zkeai/MuCoinPay/McPay-go/common/net/chttp"
	"github.com/Zkeai/MuCoinPay/McPay-go/internal/service"
)

var svc *service.Service

func InitRouter(s *chttp.Server, service *service.Service) {
	svc = service

	g := s.Group("/v1")
	ug := g.Group("/user")
	{
		ug.GET("/detail", userDetail)
	}
}
