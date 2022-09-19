package ordersfake

import (
	"time"

	"github.com/phoebetron/trades/typ/orders"
)

func (f *Fake) Update(time.Time, *orders.Orders) error {
	return nil
}
