package buffer

import (
	"time"

	"github.com/phoebetron/trades/typ/orders"
)

type Buffer interface {
	Buffer(*orders.Bundle)
	Finish(time.Time)
	Metric() int
	Orders() chan *orders.Orders
}
