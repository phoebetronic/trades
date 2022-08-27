package tradesredis

import "github.com/phoebetron/trades/typ/market"

func (r *Redis) Market() market.Interface {
	return r.mar
}
