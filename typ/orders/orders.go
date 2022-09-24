package orders

import (
	"github.com/xh3b4sd/framer"
)

func (o *Orders) Frame(c framer.Config) *Framer {
	return &Framer{
		fra: framer.New(c),
		ord: o,
	}
}
