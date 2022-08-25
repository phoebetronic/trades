package key

type Config struct {
	Exc string
	Ass string
}

// TODO rename to market ?
type Key struct {
	exc string
	ass string
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

func verify(con Config) {
	{
		if con.Exc == "" {
			panic("Exc must not be empty")
		}
		if con.Ass == "" {
			panic("Ass must not be empty")
		}
	}
}
