package tradesfake

import (
	"time"

	"github.com/phoebetronic/trades/typ/trades"
)

func (f *Fake) Create(time.Time, *trades.Trades) error {
	return nil
}
