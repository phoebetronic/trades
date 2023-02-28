package tradesredis

import (
	"github.com/phoebetronic/trades/typ/trades"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/protobuf/proto"
)

func (r *Redis) Latest() (*trades.Trade, error) {
	var err error

	var key string
	{
		key = r.Key()
	}

	var val string
	{
		res, err := r.sor.Search().Order(key, 0, 0)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if len(res) == 0 {
			return nil, tracer.Maskf(notFoundError, "latest trade does not exist")
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
			return nil, tracer.Maskf(notFoundError, "latest trade does not exist")
		}
	}

	return tra.TR[len(tra.TR)-1], nil
}
