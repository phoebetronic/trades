package trades

import (
	"time"

	"github.com/phoebetron/trades/typ/key"
)

type Storage interface {
	Create(time.Time, *Trades) error
	Delete(time.Time) error
	Latest() (*Trade, error)
	Market() key.Interface
	Rigmos() (*Trade, error)
	Search(time.Time) (*Trades, error)
	Update(time.Time, *Trades) error
}
