package tradesfake

type Config struct{}

type Fake struct{}

func New(Config) (*Fake, error) {
	var f *Fake
	{
		f = &Fake{}
	}

	return f, nil
}
