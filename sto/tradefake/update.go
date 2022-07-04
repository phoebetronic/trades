package fake

import (
	"time"

	"github.com/phoebetron/trades/typ/trade"
)

func (f *Fake) Update(time.Time, []trade.Trade) error {
	return nil
}
