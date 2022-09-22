package trades

import (
	"time"

	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Framer struct {
	dur time.Duration
	fra *framer.Framer
	his []*Trades
	ind int
	las *Trade
	pre int
	tra *Trades
}

func (f *Framer) Hist() []*Trades {
	return f.his
}

func (f *Framer) Last() bool {
	return f.fra.Last() || len(f.tra.TR) == 0
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
			TS: timestamppb.New(f.las.TS.AsTime().Truncate(f.dur).Add(f.dur)),
		}}
	}

	{
		f.las = tra.LA()
	}

	{
		f.his = f.hist(tra)
	}

	return tra
}

func (f *Framer) Pred() int {
	return f.pre
}

func (f *Framer) hist(tra *Trades) []*Trades {
	if !f.his[len(f.his)-1].ST.AsTime().Add(time.Minute).After(tra.ST.AsTime()) {
		f.his = append(f.his, tra)
	}

	if len(f.his) > 5 {
		{
			copy(f.his[0:], f.his[1:])
			f.his[len(f.his)-1] = nil
			f.his = f.his[:len(f.his)-1]
		}

		{
			f.pre = 0
		}

		for i := 1; i < 5; i++ {
			f.pre += f.pred(f.his[i-1].PR().Avg(), f.his[i].PR().Avg())
		}
	}

	return f.his
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

	for ; f.ind < len(f.tra.TR); f.ind++ {
		if !f.tra.TR[f.ind].TS.AsTime().Before(fra.End) {
			break
		}
		if f.tra.TR[f.ind].TS.AsTime().Before(fra.Sta) {
			continue
		}

		{
			tra.TR = append(tra.TR, f.tra.TR[f.ind])
		}
	}

	// Since the frame creation moves the trades window forward, we push the
	// start time of the source structure to the end time of the constructured
	// frame, because there are no trades for the constructured frame range
	// anymore within the source structure.
	{
		f.tra.ST = timestamppb.New(fra.End)
	}

	return tra
}

func (f *Framer) pred(a float32, b float32) int {
	if a > b {
		return -1
	}

	if a < b {
		return +1
	}

	return 0
}
