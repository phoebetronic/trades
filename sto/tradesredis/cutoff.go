package tradesredis

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Cutoff(lim int) error {
	var key string
	{
		key = r.Key()
	}

	{
		err := r.sor.Delete().Limit(key, lim)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
