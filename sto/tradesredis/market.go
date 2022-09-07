package tradesredis

import "github.com/phoebetron/trades/typ/market"

func (r *Redis) Market() market.Market {
	return r.mar
}
