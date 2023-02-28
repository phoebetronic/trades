package ordersfake

import (
	"time"

	"github.com/phoebetronic/trades/typ/orders"
)

func (f *Fake) Search(time.Time) (*orders.Orders, error) {
	return nil, nil
}
