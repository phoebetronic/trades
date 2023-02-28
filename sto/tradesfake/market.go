package tradesfake

import (
	"github.com/phoebetronic/trades/typ/market"
)

func (f *Fake) Market() market.Market {
	return &market.M{}
}
