package fake

import "github.com/phoebetron/trades/typ/trade"

func (f *Fake) Latest() (trade.Trade, error) {
	return trade.Trade{}, nil
}
