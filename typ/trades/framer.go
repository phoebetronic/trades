package trades

import (
	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Framer struct {
	fra *framer.Framer
	tra *Trades
}

func (f *Framer) Last() bool {
	return f.fra.Last() || len(f.tra.TR) == 0
}

func (f *Framer) Next() *Trades {
	if f.Last() {
		return nil
	}

	var tra *Trades
	for {
		var fra framer.Frame
		{
			fra = f.fra.Next()
		}

		{
			tra = f.next(fra)
		}

		if len(tra.TR) != 0 {
			break
		}
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

	var ind int
	for i := 0; i < len(f.tra.TR); i++ {
		{
			ind = i
		}

		if !f.tra.TR[i].TS.AsTime().Before(fra.End) {
			break
		}
		if f.tra.TR[i].TS.AsTime().Before(fra.Sta) {
			continue
		}

		{
			tra.TR = append(tra.TR, f.tra.TR[i])
		}
	}

	// Once the next frame of trades got constructed, we remove the allocated
	// trades from the source structure. This is to move trades around instead
	// of duplicating them.
	{
		copy(f.tra.TR[0:], f.tra.TR[ind:])
		for k, n := len(f.tra.TR)-ind, len(f.tra.TR); k < n; k++ {
			f.tra.TR[k] = nil
		}
		f.tra.TR = f.tra.TR[:len(f.tra.TR)-ind]
	}

	// In case there is a last remaining trade, we simply empty the internal
	// list of trades, because there are no frames left for iteration.
	if len(f.tra.TR) == 1 {
		f.tra.TR = nil
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
