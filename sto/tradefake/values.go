package fake

import (
	"time"

	"github.com/phoebetron/trades/typ/trade"
)

func (f *Fake) Values() (map[time.Time][]trade.Trade, error) {
	return nil, nil
}
