package market

import "time"

type M struct {
	exc string
	ass string
	dur time.Duration
}

func New(con Config) *M {
	{
		con.Verify()
	}

	return &M{
		exc: con.Exc,
		ass: con.Ass,
		dur: con.Dur,
	}
}

func (m *M) Exc() string {
	return m.exc
}

func (m *M) Ass() string {
	return m.ass
}

func (m *M) Dur() time.Duration {
	return m.dur
}
