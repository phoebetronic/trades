package buffer

import (
	"time"

	"github.com/phoebetronic/trades/typ/trades"
)

type Buffer interface {
	Buffer(*trades.Trade)
	Finish(time.Time)
	Latest(*trades.Trade)
	Metric() int
	Trades() chan *trades.Trades
}
