package tradesredis

import (
	"github.com/phoebetronic/trades/typ/market"
	"github.com/xh3b4sd/redigo/pkg/sorted"
)

type Config struct {
	Mar market.Market
	Sor sorted.Interface
}

func (c Config) Verify() {
	if c.Sor == nil {
		panic("Config.Sor must not be empty")
	}
	if c.Mar == nil {
		panic("Config.Mar must not be empty")
	}
}
