package tradesfake

import "github.com/phoebetron/trades/typ/key"

func (f *Fake) Market() key.Interface {
	return &key.Key{}
}
