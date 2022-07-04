package traderedis

import (
	"encoding/json"
	"time"

	"github.com/xh3b4sd/redigo/pkg/sorted"
	"github.com/xh3b4sd/tracer"

	"github.com/phoebetron/trades/typ/trade"
)

func (r *Redis) Update(day time.Time, can []trade.Trade) error {
	var key string
	{
		key = r.Key()
	}

	var val string
	{
		byt, err := json.Marshal(can)
		if err != nil {
			return tracer.Mask(err)
		}

		val = string(byt)
	}

	var sco float64
	{
		sco = float64(day.Unix())
	}

	{
		_, err := r.sor.Update().Value(key, val, sco)
		if sorted.IsNotFound(err) {
			return tracer.Maskf(notFoundError, "candle for %s does not exist", day.String())
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
