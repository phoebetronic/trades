package traderedis

import "time"

const (
	trafmt = "06-01-02"
)

func timday(tim time.Time) time.Time {
	return tim.UTC().Truncate(24 * time.Hour)
}

func timfmt(tim time.Time) string {
	return tim.UTC().Format(trafmt)
}
