package traderedis

import (
	"encoding/json"
	"time"

	"github.com/xh3b4sd/tracer"

	"github.com/phoebetron/trades/typ/trade"
)

func (r *Redis) Values() (map[time.Time][]trade.Trade, error) {
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

	tra := map[time.Time][]trade.Trade{}
	{
		for _, r := range res {
			var c []trade.Trade
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
