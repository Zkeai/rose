package main

import (
	"context"
	"flag"
	"github.com/Zkeai/MuCoinPay/McPay-go/configs"
	"github.com/Zkeai/MuCoinPay/McPay-go/pkg/db"
	"github.com/Zkeai/MuCoinPay/McPay-go/pkg/logger"
	"github.com/Zkeai/MuCoinPay/McPay-go/pkg/redis"
	"os"
	"os/signal"
	"syscall"

	cconf "github.com/Zkeai/MuCoinPay/McPay-go/common/conf"
	"github.com/Zkeai/MuCoinPay/McPay-go/internal/conf"
	"github.com/Zkeai/MuCoinPay/McPay-go/internal/server"
)

var filePath = flag.String("conf", "etc/config.yaml", "the config path")

func main() {
	flag.Parse()
	configs.Setup()
	logger.Setup()
	db.ConnectMySQL()
	redis.ConnectRedis()

	c := new(conf.Conf)
	if err := cconf.Unmarshal(*filePath, c); err != nil {
		panic(err)
	}
	srv := server.NewHTTP(c)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_ = srv.Shutdown(context.Background())
			return
		default:
			return
		}
	}
}
