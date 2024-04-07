package service

import (
	"github.com/Zkeai/MuCoinPay/McPay-go/internal/conf"
)

type Service struct {
	conf *conf.Conf
}

func NewService(conf *conf.Conf) *Service {
	return &Service{
		conf: conf,
	}
}
