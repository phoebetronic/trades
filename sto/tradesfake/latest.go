package tradesfake

import "github.com/phoebetronic/trades/typ/trades"

func (f *Fake) Latest() (*trades.Trade, error) {
	return nil, nil
}
