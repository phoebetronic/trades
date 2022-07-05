package tradesredis

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/redigo/pkg/sorted"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/proto"
)

func (r *Redis) Update(tim time.Time, tra *trades.Trades) error {
	var key string
	{
		key = r.Key()
	}

	var val string
	{
		byt, err := proto.Marshal(tra)
		if err != nil {
			return tracer.Mask(err)
		}

		val = string(byt)
	}

	var sco float64
	{
		sco = float64(tim.Unix())
	}

	{
		_, err := r.sor.Update().Value(key, val, sco)
		if sorted.IsNotFound(err) {
			return tracer.Maskf(notFoundError, "trade for %s does not exist", tim.String())
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
