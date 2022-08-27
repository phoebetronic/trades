package market

import "time"

type Config struct {
	Exc string
	Ass string
	Dur time.Duration
}

func (c Config) Verify() {
	if c.Exc == "" {
		panic("Config.Exc must not be empty")
	}
	if c.Ass == "" {
		panic("Config.Ass must not be empty")
	}
	if c.Dur == 0 {
		panic("Config.Dur must not be empty")
	}
}
