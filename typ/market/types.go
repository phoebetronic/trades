package market

import "time"

type Interface interface {
	Exc() string
	Ass() string
	Dur() time.Duration
}
