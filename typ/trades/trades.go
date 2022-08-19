package trades

import (
	"sort"

	"github.com/phoebetron/trades/typ/floats"
	"github.com/xh3b4sd/framer"
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

func (t *Trades) Frame(c framer.Config) *Framer {
	return &Framer{
		fra: framer.New(c),
		tra: t,
	}
}

func (t *Trades) Merge(l []*Trades) *Trades {
	if len(l) == 0 {
		return t
	}

	{
		sort.SliceStable(l, func(i, j int) bool { return l[i].ST.AsTime().Unix() < l[j].ST.AsTime().Unix() })
	}

	var i int
	{
		i = len(l) - 1
	}

	{
		t.EX = l[0].EX
		t.AS = l[0].AS
		t.ST = l[0].ST
		t.EN = l[i].EN
	}

	for _, v := range l {
		t.TR = append(t.TR, v.TR...)
	}

	return t
}

func (t *Trades) Scale(f float32) *Trades {
	var tr []*Trade

	for _, v := range t.TR {
		tr = append(tr, v.Scale(f))
	}

	return t.cop(tr)
}

func (t *Trades) cop(l []*Trade) *Trades {
	return &Trades{
		EX: t.EX,
		AS: t.AS,
		ST: tim(t.ST),
		EN: tim(t.EN),
		TR: l,
	}
}

func tim(t *timestamppb.Timestamp) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(t.AsTime())
}
