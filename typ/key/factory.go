package key

func Default() *Key {
	var err error

	var sto *Key
	{
		c := Config{
			Exc: "ftx",
			Ass: "eth",
		}

		sto, err = New(c)
		if err != nil {
			panic(err)
		}
	}

	return sto
}
