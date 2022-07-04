package traderedis

import (
	"encoding/json"
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Values() (map[time.Time][]trades.Trade, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var res []string
	{
		res, err = r.sor.Search().Order(key, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	tra := map[time.Time][]trades.Trade{}
	{
		for _, r := range res {
			var c []trades.Trade
			{
				err = json.Unmarshal([]byte(r), &c)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				if len(c) != 0 {
					tra[timday(c[0].TS)] = c
				}
			}
		}
	}

	return tra, nil
}
