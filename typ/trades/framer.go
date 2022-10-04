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
	tra *Trades
}

func (f *Framer) Conf() framer.Config {
	return f.con
}

func (f *Framer) Rang(off time.Duration, win time.Duration, rig time.Duration, tic time.Duration) *Ranger {
	var o time.Time
	var w time.Time
	var s time.Time
	var r time.Time
	{
		o = f.cur.Sta.Add(-off).Add(-win)
		w = f.cur.Sta.Add(-win)
		s = f.cur.Sta
		r = f.cur.Sta.Add(+rig)
	}

	return &Ranger{
		off: f.lis(framer.Config{Sta: o, End: w, Len: tic}),
		win: f.lis(framer.Config{Sta: w, End: s, Len: tic}),
		rig: f.lis(framer.Config{Sta: s, End: r, Len: tic}),
	}
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

func (f *Framer) lis(con framer.Config) []*Trades {
	var fir *Trade
	for i := f.ind; i >= 0; i-- {
		if !f.tra.TR[i].TS.AsTime().After(con.Sta) {
			fir = f.tra.TR[i]
			break
		}
	}

	if fir == nil {
		fir = f.tra.TR[0]
	}

	var fra *Framer
	{
		fra = &Framer{
			con: con,
			fra: framer.New(con),
			las: fir,
			tra: f.tra,
		}
	}

	var tra []*Trades
	for !fra.Last() {
		tra = append(tra, fra.Next())
	}

	return tra
}
