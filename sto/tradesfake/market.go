package tradesfake

import (
	"github.com/phoebetron/trades/typ/market"
)

func (f *Fake) Market() market.Interface {
	return &market.Market{}
}
