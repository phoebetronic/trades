package tradesfake

import (
	"time"

	"github.com/phoebetronic/trades/typ/trades"
)

func (f *Fake) Update(time.Time, *trades.Trades) error {
	return nil
}
