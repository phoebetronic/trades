package tradesredis

import (
	"fmt"
)

func (r *Redis) Key() string {
	return fmt.Sprintf("tra:raw|exc:%s|ass:%s", r.mar.Exc(), r.mar.Ass())
}
