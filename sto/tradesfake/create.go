package tradesfake

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

func (f *Fake) Create(time.Time, *trades.Trades) error {
	return nil
}
