package traderedis

import (
	"encoding/json"
	"time"

	"github.com/xh3b4sd/tracer"

	"github.com/phoebetron/trades/typ/trade"
)

func (r *Redis) Search(day time.Time) ([]trade.Trade, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var sco float64
	{
		sco = float64(day.Unix())
	}

	var val string
	{
		res, err := r.sor.Search().Score(key, sco, sco)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if len(res) == 0 {
			return nil, tracer.Maskf(notFoundError, "candles for %s do not exist", day.String())
		}

		val = res[0]
	}

	var can []trade.Trade
	{
		err = json.Unmarshal([]byte(val), &can)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return can, nil
}
