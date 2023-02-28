package ordersredis

import (
	"github.com/phoebetronic/trades/typ/market"
	"github.com/xh3b4sd/redigo/pkg/sorted"
)

type Redis struct {
	mar market.Market
	sor sorted.Interface
}

func New(con Config) *Redis {
	{
		con.Verify()
	}

	return &Redis{
		mar: con.Mar,
		sor: con.Sor,
	}
}
