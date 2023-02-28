package tradesfake

import (
	"time"

	"github.com/phoebetronic/trades/typ/trades"
)

func (f *Fake) Search(time.Time) (*trades.Trades, error) {
	return nil, nil
}
