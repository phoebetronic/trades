package trades

import "time"

type Storage interface {
	Create(time.Time, *Trades) error
	Delete(time.Time) error
	Latest() (*Trade, error)
	Search(time.Time) (*Trades, error)
	Update(time.Time, *Trades) error
}
