package orders

import (
	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Framer struct {
	fra *framer.Framer
	ord *Orders
}

func (f *Framer) Last() bool {
	return f.fra.Last() || len(f.ord.BU) == 0
}

func (f *Framer) Next() *Orders {
	if f.Last() {
		return nil
	}

	var fra framer.Frame
	{
		fra = f.fra.Next()
	}

	var ord *Orders
	{
		ord = f.next(fra)
	}

	return ord
}

func (f *Framer) next(fra framer.Frame) *Orders {
	var ord *Orders
	{
		ord = &Orders{
			EX: f.ord.EX,
			AS: f.ord.AS,
			ST: timestamppb.New(fra.Sta),
			EN: timestamppb.New(fra.End),
		}
	}

	var ind int
	for i := 0; i < len(f.ord.BU); i++ {
		{
			ind = i
		}

		if !f.ord.BU[i].TS.AsTime().Before(fra.End) {
			break
		}
		if f.ord.BU[i].TS.AsTime().Before(fra.Sta) {
			continue
		}

		{
			ord.BU = append(ord.BU, f.ord.BU[i])
		}
	}

	// Once the next frame of trades got constructed, we remove the allocated
	// trades from the source structure. This is to move trades around instead
	// of duplicating them.
	{
		copy(f.ord.BU[0:], f.ord.BU[ind:])
		for k, n := len(f.ord.BU)-ind, len(f.ord.BU); k < n; k++ {
			f.ord.BU[k] = nil
		}
		f.ord.BU = f.ord.BU[:len(f.ord.BU)-ind]
	}

	// In case there is a last remaining trade, we simply empty the internal
	// list of trades, because there are no frames left for iteration.
	if len(f.ord.BU) == 1 && f.fra.Last() {
		f.ord.BU = nil
	}

	// Since the frame creation moves the trades window forward, we push the
	// start time of the source structure to the end time of the constructured
	// frame, because there are no trades for the constructured frame range
	// anymore within the source structure.
	{
		f.ord.ST = timestamppb.New(fra.End)
	}

	return ord
}
