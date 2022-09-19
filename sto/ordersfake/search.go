package ordersfake

import (
	"time"

	"github.com/phoebetron/trades/typ/orders"
)

func (f *Fake) Search(time.Time) (*orders.Orders, error) {
	return nil, nil
}
