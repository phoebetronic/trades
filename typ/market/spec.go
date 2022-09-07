package market

import "time"

type Market interface {
	Exc() string
	Ass() string
	Dur() time.Duration
}
