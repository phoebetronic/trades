package tradesredis

import (
	"github.com/phoebetron/trades/typ/key"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/redigo"
)

func Default() trades.Storage {
	var err error

	var red *Redis
	{
		c := Config{
			Key: key.Default(),
			Sor: redigo.Default().Sorted(),
		}

		red, err = New(c)
		if err != nil {
			panic(err)
		}
	}

	return red
}
