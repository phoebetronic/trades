package tradesfake

import "github.com/phoebetron/trades/typ/trades"

func (f *Fake) Latest() (trades.Trade, error) {
	return trades.Trade{}, nil
}
