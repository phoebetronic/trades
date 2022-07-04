package tradesredis

import (
	"encoding/json"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Latest() (trades.Trade, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var val string
	{
		res, err := r.sor.Search().Order(key, 0, 1)
		if err != nil {
			return trades.Trade{}, tracer.Mask(err)
		}

		if len(res) == 0 {
			return trades.Trade{}, tracer.Maskf(notFoundError, "latest trade does not exist")
		}
		if len(res) != 1 {
			return trades.Trade{}, tracer.Maskf(executionFailedError, "unexpected redis response")
		}

		val = res[0]
	}

	var tra []trades.Trade
	{
		err = json.Unmarshal([]byte(val), &tra)
		if err != nil {
			return trades.Trade{}, tracer.Mask(err)
		}

		if len(tra) == 0 {
			return trades.Trade{}, tracer.Maskf(notFoundError, "latest trade does not exist")
		}
	}

	return tra[len(tra)-1], nil
}
