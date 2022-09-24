package trades

import (
	"time"

	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Framer struct {
	con framer.Config
	cur framer.Frame
	fra *framer.Framer
	ind int
	las *Trade
	pre int
	tra *Trades
}

func (f *Framer) Hist(len time.Duration, tic time.Duration) []*Trades {
	var c framer.Config
	{
		c = framer.Config{
			Sta: f.cur.Sta.Add(-len),
			End: f.cur.Sta,
			Len: tic,
		}
	}

	var fir *Trade
	for i := f.ind; i >= 0; i-- {
		if !f.tra.TR[i].TS.AsTime().After(c.Sta) {
			fir = f.tra.TR[i]
			break
		}
	}

	if fir == nil {
		fir = f.tra.TR[0]
	}

	var n *Framer
	{
		n = &Framer{
			con: c,
			fra: framer.New(c),
			las: fir,
			tra: f.tra,
		}
	}

	var t []*Trades
	for !n.Last() {
		t = append(t, n.Next())
	}

	return t
}

func (f *Framer) Last() bool {
	return f.fra.Last()
}

func (f *Framer) Next() *Trades {
	if f.Last() {
		return nil
	}

	var fra framer.Frame
	{
		fra = f.fra.Next()
	}

	var tra *Trades
	{
		tra = f.next(fra)
	}

	if len(tra.TR) == 0 {
		tra.TR = []*Trade{{
			LI: f.las.LI,
			PR: f.las.PR,
			LO: f.las.LO,
			SH: f.las.SH,
			TS: timestamppb.New(fra.Sta),
		}}
	}

	{
		f.cur = fra
		f.las = tra.LA()
	}

	return tra
}

func (f *Framer) next(fra framer.Frame) *Trades {
	var tra *Trades
	{
		tra = &Trades{
			EX: f.tra.EX,
			AS: f.tra.AS,
			ST: timestamppb.New(fra.Sta),
			EN: timestamppb.New(fra.End),
		}
	}

	for i := f.ind; i < len(f.tra.TR); i++ {
		if f.tra.TR[i].TS.AsTime().Before(fra.Sta) {
			f.ind++
			continue
		}
		if !f.tra.TR[i].TS.AsTime().Before(fra.End) {
			break
		}

		{
			tra.TR = append(tra.TR, f.tra.TR[i])
		}
	}

	return tra
}
