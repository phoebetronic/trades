package tradesfake

func Default() *Fake {
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
