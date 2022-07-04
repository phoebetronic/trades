package tradesredis

import (
	"encoding/json"
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Search(day time.Time) ([]trades.Trade, error) {
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
			return nil, tracer.Maskf(notFoundError, "trades for %s do not exist", day.String())
		}

		val = res[0]
	}

	var tra []trades.Trade
	{
		err = json.Unmarshal([]byte(val), &tra)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return tra, nil
}
