package tradesfake

import "github.com/phoebetron/trades/typ/trades"

func Default() trades.Storage {
	var err error

	var fak *Fake
	{
		c := Config{}

		fak, err = New(c)
		if err != nil {
			panic(err)
		}
	}

	return fak
}
