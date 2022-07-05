package tradesredis

import "github.com/phoebetron/trades/typ/key"

func (r *Redis) Market() key.Interface {
	return r.key
}
