package orders

import (
	"time"

	"github.com/phoebetronic/trades/typ/market"
)

type Storage interface {
	Create(time.Time, *Orders) error
	Delete(time.Time) error
	Market() market.Market
	Search(time.Time) (*Orders, error)
	Update(time.Time, *Orders) error
}
