package conf

import (
	"github.com/Zkeai/MuCoinPay/McPay-go/common/net/chttp"
)

type Conf struct {
	Server *chttp.Config `yaml:"server"`
}
