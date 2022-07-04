package traderedis

import (
	"github.com/xh3b4sd/redigo"

	"github.com/phoebetron/trades/typ/key"
)

func Default() *Redis {
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
