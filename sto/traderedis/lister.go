package traderedis

import (
	"encoding/json"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Lister() ([]string, error) {
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

	var lis []string
	{
		for _, r := range res {
			var c []trades.Trade
			err = json.Unmarshal([]byte(r), &c)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if len(c) != 0 {
				lis = append(lis, timfmt(c[0].TS))
			}
		}
	}

	return lis, nil
}
