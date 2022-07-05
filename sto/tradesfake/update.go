package tradesfake

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

func (f *Fake) Update(time.Time, *trades.Trades) error {
	return nil
}
