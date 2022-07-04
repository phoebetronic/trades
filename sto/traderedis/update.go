package traderedis

import (
	"encoding/json"
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/redigo/pkg/sorted"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Update(day time.Time, tra []trades.Trade) error {
	var key string
	{
		key = r.Key()
	}

	var val string
	{
		byt, err := json.Marshal(tra)
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
			return tracer.Maskf(notFoundError, "trade for %s does not exist", day.String())
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
