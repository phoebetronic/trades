package ordersredis

import (
	"time"

	"github.com/phoebetronic/trades/typ/orders"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/proto"
)

func (r *Redis) Search(tim time.Time) (*orders.Orders, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var sco float64
	{
		sco = float64(tim.Unix())
	}

	var val string
	{
		res, err := r.sor.Search().Score(key, sco, sco)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if len(res) == 0 {
			return nil, tracer.Maskf(notFoundError, "orders for %s do not exist", tim.String())
		}

		val = res[0]
	}

	ord := &orders.Orders{}
	{
		err = proto.Unmarshal([]byte(val), ord)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return ord, nil
}
