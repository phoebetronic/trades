package buffer

import (
	"time"

	"github.com/phoebetron/trades/typ/key"
)

type Config struct {
	Len time.Duration
	Mar *key.Key
}

func (c Config) verify() {
	if c.Len == 0 {
		panic("Config.Len must not be empty")
	}
	if c.Mar == nil {
		panic("Config.Mar must not be empty")
	}
}
