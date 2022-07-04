package traderedis

import (
	"time"

	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(day time.Time) error {
	var key string
	{
		key = r.Key()
	}

	var sco float64
	{
		sco = float64(day.Unix())
	}

	{
		err := r.sor.Delete().Score(key, sco)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
