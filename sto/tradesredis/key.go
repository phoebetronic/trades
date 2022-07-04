package tradesredis

import (
	"fmt"
)

func (r *Redis) Key() string {
	return fmt.Sprintf("tra:raw|exc:%s|ass:%s", r.key.Exc(), r.key.Ass())
}
