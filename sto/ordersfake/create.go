package ordersfake

import (
	"time"

	"github.com/phoebetronic/trades/typ/orders"
)

func (f *Fake) Create(time.Time, *orders.Orders) error {
	return nil
}
