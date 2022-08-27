package market

import "time"

type Market struct {
	exc string
	ass string
	dur time.Duration
}

func New(con Config) *Market {
	{
		con.Verify()
	}

	return &Market{
		exc: con.Exc,
		ass: con.Ass,
		dur: con.Dur,
	}
}

func (m *Market) Exc() string {
	return m.exc
}

func (m *Market) Ass() string {
	return m.ass
}

func (m *Market) Dur() time.Duration {
	return m.dur
}
