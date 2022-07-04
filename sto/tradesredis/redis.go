package tradesredis

import (
	"github.com/phoebetron/trades/typ/key"
	"github.com/xh3b4sd/redigo/pkg/sorted"
)

type Config struct {
	Key key.Interface
	Sor sorted.Interface
}

type Redis struct {
	key key.Interface
	sor sorted.Interface
}

func New(con Config) (*Redis, error) {
	{
		verify(con)
	}

	var r *Redis
	{
		r = &Redis{
			key: con.Key,
			sor: con.Sor,
		}
	}

	return r, nil
}

func verify(con Config) {
	{
		if con.Sor == nil {
			panic("Sor must not be empty")
		}
		if con.Key == nil {
			panic("Key must not be empty")
		}
	}
}
