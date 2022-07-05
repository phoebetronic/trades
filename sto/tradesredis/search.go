package tradesredis

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/proto"
)

func (r *Redis) Search(tim time.Time) (*trades.Trades, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var sco float64
	{
		sco = float64(tim.Unix())
	}

	var val string
	{
		res, err := r.sor.Search().Score(key, sco, sco)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if len(res) == 0 {
			return nil, tracer.Maskf(notFoundError, "trades for %s do not exist", tim.String())
		}

		val = res[0]
	}

	var tra *trades.Trades
	{
		err = proto.Unmarshal([]byte(val), tra)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return tra, nil
}
