package ordersfake

import (
	"github.com/phoebetronic/trades/typ/market"
)

func (f *Fake) Market() market.Market {
	return &market.M{}
}
