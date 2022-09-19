package ordersredis

import (
	"time"

	"github.com/phoebetron/trades/typ/orders"
	"github.com/xh3b4sd/redigo/pkg/sorted"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/proto"
)

func (r *Redis) Create(day time.Time, ord *orders.Orders) error {
	var key string
	{
		key = r.Key()
	}

	var val string
	{
		byt, err := proto.Marshal(ord)
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
		err := r.sor.Create().Element(key, val, sco)
		if sorted.IsAlreadyExistsError(err) {
			return tracer.Mask(alreadyExistsError)
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
