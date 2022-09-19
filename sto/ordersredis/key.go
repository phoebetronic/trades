package ordersredis

import (
	"fmt"
)

func (r *Redis) Key() string {
	return fmt.Sprintf("ord:raw|exc:%s|ass:%s", r.mar.Exc(), r.mar.Ass())
}
