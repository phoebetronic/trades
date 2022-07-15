package trades

import "google.golang.org/protobuf/types/known/timestamppb"

func (t *Trade) Empty() bool {
	return t.PR == 0 && t.LO == 0 && t.SH == 0 && t.TS.AsTime().IsZero()
}

func (t *Trade) Scale(f float32) *Trade {
	return &Trade{
		LI: t.LI,
		PR: t.PR + (t.PR * f),
		LO: t.LO + (t.LO * f),
		SH: t.SH + (t.SH * f),
		TS: timestamppb.New(t.TS.AsTime()),
	}
}
