package key

import (
	"time"
)

type Config struct {
	Exc string
	Ass string
	Res time.Duration
}

type Key struct {
	exc string
	ass string
	res time.Duration
}

func New(con Config) (*Key, error) {
	{
		verify(con)
	}

	var k *Key
	{
		k = &Key{
			exc: con.Exc,
			ass: con.Ass,
			res: con.Res,
		}
	}

	return k, nil
}

func (k *Key) Exc() string {
	return k.exc
}

func (k *Key) Ass() string {
	return k.ass
}

func (k *Key) Res() string {
	return k.res.String()
}

func verify(con Config) {
	{
		if con.Exc == "" {
			panic("Exc must not be empty")
		}
		if con.Ass == "" {
			panic("Ass must not be empty")
		}
		if con.Res == 0 {
			panic("Res must not be empty")
		}
	}
}
