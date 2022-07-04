package fake

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

func (f *Fake) Update(time.Time, []trades.Trade) error {
	return nil
}
