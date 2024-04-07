package server

import (
	"github.com/Zkeai/rose/common/net/chttp"
	"github.com/Zkeai/rose/internal/conf"
	"github.com/Zkeai/rose/internal/handler"
	"github.com/Zkeai/rose/internal/service"
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
