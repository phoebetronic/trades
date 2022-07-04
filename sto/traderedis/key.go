package traderedis

import (
	"fmt"
)

func (r *Redis) Key() string {
	return fmt.Sprintf("typ:tra|exc:%s|ass:%s|res:%s", r.key.Exc(), r.key.Ass(), r.key.Res())
}
