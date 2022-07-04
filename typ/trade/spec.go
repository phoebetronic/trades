package trade

import "time"

type Storage interface {
	Create(time.Time, []Trade) error
	Cutoff(lim int) error
	Delete(time.Time) error
	Latest() (Trade, error)
	Lister() ([]string, error)
	Search(time.Time) ([]Trade, error)
	Update(time.Time, []Trade) error
	Values() (map[time.Time][]Trade, error)
}
