package traderedis

import (
	"encoding/json"

	"github.com/xh3b4sd/tracer"

	"github.com/phoebetron/trades/typ/trade"
)

func (r *Redis) Latest() (trade.Trade, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var val string
	{
		res, err := r.sor.Search().Order(key, 0, 1)
		if err != nil {
			return trade.Trade{}, tracer.Mask(err)
		}

		if len(res) == 0 {
			return trade.Trade{}, tracer.Maskf(notFoundError, "latest trade does not exist")
		}
		if len(res) != 1 {
			return trade.Trade{}, tracer.Maskf(executionFailedError, "unexpected redis response")
		}

		val = res[0]
	}

	var tra []trade.Trade
	{
		err = json.Unmarshal([]byte(val), &tra)
		if err != nil {
			return trade.Trade{}, tracer.Mask(err)
		}

		if len(tra) == 0 {
			return trade.Trade{}, tracer.Maskf(notFoundError, "latest trade does not exist")
		}
	}

	return tra[len(tra)-1], nil
}
