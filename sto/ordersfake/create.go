package ordersfake

import (
	"time"

	"github.com/phoebetron/trades/typ/orders"
)

func (f *Fake) Create(time.Time, *orders.Orders) error {
	return nil
}
