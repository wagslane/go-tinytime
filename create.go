package tinytime

import (
	"time"
)

// New creates a new TinyTime, similar to the time.Time package,
// months are specifed from 1-12, and days are specified 1-31, depending on the month
func New(unixSeconds uint32) TinyTime {
	return TinyTime{
		unix: unixSeconds,
	}
}

// Now returns the current time
func Now() TinyTime {
	tt, _ := FromTime(time.Now())
	return tt
}

// Unix creates a tinytime from seconds and nanoseconds. As usual, this is truncated to
// the nearest day
func Unix(sec int64, nsec int64) TinyTime {
	tt, _ := FromTime(time.Unix(sec, nsec).UTC())
	return tt
}
