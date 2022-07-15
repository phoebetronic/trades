package trades

import (
	"github.com/phoebetron/trades/typ/floats"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *Trades) PR() floats.Floats {
	var f floats.Floats

	for _, t := range t.TR {
		f.FL = append(f.FL, t.PR)
	}

	return f
}

func (t *Trades) LO() floats.Floats {
	var f floats.Floats

	for _, t := range t.TR {
		f.FL = append(f.FL, t.LO)
	}

	return f
}

func (t *Trades) SH() floats.Floats {
	var f floats.Floats

	for _, t := range t.TR {
		f.FL = append(f.FL, t.SH)
	}

	return f
}

func (t *Trades) Scale(f float32) *Trades {
	var tr []*Trade

	for _, v := range t.TR {
		tr = append(tr, v.Scale(f))
	}

	return t.cop(tr)
}

func (t *Trades) cop(tr []*Trade) *Trades {
	return &Trades{
		EX: t.EX,
		AS: t.AS,
		ST: copTim(t.ST),
		EN: copTim(t.EN),
		TR: tr,
	}
}

func copTim(tim *timestamppb.Timestamp) *timestamppb.Timestamp {
	if tim == nil {
		return nil
	}

	return timestamppb.New(tim.AsTime())
}
