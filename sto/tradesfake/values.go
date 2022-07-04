package tradesfake

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

func (f *Fake) Values() (map[time.Time][]trades.Trade, error) {
	return nil, nil
}
