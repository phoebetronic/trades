package tradesredis

import "github.com/phoebetronic/trades/typ/market"

func (r *Redis) Market() market.Market {
	return r.mar
}
