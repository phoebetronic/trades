package buffer

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

type Interface interface {
	Buffer(*trades.Trade)
	Finish(time.Time)
	Metric() (int, int, int)
	Trades() chan *trades.Trades
}
