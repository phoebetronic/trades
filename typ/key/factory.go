package key

import "time"

func Default() *Key {
	var err error

	var sto *Key
	{
		c := Config{
			Exc: "ftx",
			Ass: "eth",
			Res: 60 * time.Second,
		}

		sto, err = New(c)
		if err != nil {
			panic(err)
		}
	}

	return sto
}
