package tradesredis

import (
	"github.com/phoebetron/trades/typ/trades"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/proto"
)

func (r *Redis) Rigmos() (*trades.Trade, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var val string
	{
		res, err := r.sor.Search().Order(key, -1, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if len(res) == 0 {
			return nil, tracer.Maskf(notFoundError, "first trade does not exist")
		}
		if len(res) != 1 {
			return nil, tracer.Maskf(executionFailedError, "unexpected redis response")
		}

		val = res[0]
	}

	tra := &trades.Trades{}
	{
		err = proto.Unmarshal([]byte(val), tra)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if len(tra.TR) == 0 {
			return nil, tracer.Maskf(notFoundError, "first trade does not exist")
		}
	}

	return tra.TR[0], nil
}
