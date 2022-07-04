package trades

import (
	"time"
)

type Trade struct {
	// LI expresses whether this particular trade was caused by a liquidation.
	LI bool `json:"li,omitempty"`
	// PR is the price at which this trade was made.
	PR float64 `json:"pr"`
	// LO is the size of a long trade, if it was not a short trade.
	LO float64 `json:"lo,omitempty"`
	// SH is the size of a short trade, if it was not a long trade.
	SH float64 `json:"sh,omitempty"`
	// TS is the timestamp at which this particular trade happened.
	TS time.Time `json:"ts"`
}

func (t Trade) Empty() bool {
	return t.PR == 0 && t.LO == 0 && t.SH == 0 && t.TS.IsZero()
}
