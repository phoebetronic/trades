package ordersfake

import (
	"github.com/phoebetron/trades/typ/market"
)

func (f *Fake) Market() market.Market {
	return &market.M{}
}
