package tinytime

import "time"

// After returns true if tt is after tu
func (tt TinyTime) After(tu TinyTime) bool {
	return tt.unix > tu.unix
}

// Before returns true if tt is before tu
func (tt TinyTime) Before(tu TinyTime) bool {
	return tt.unix < tu.unix
}

// Add a duration to a TinyTime. Will only have an effect if more than 1 day is added
func (tt TinyTime) Add(d time.Duration) TinyTime {
	newTD, _ := FromTime(tt.ToTime().Add(d))
	return newTD
}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
func (tt TinyTime) AddDate(years int, months int, days int) TinyTime {
	newTD, _ := FromTime(tt.ToTime().AddDate(years, months, days))
	return newTD
}

// Equal returns true if the times are exactly the same
func (tt TinyTime) Equal(tu TinyTime) bool {
	return tt.unix == tu.unix
}

// Sub returns the duration (in days) between tt and tu
func (tt TinyTime) Sub(tu TinyTime) time.Duration {
	return tt.ToTime().Sub(tu.ToTime())
}
