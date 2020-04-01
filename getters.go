package tinytime

import "time"

// Year returns the year as the readable numerical value e.g. 2020
func (tt TinyTime) Year() int {
	tm := tt.ToTime()
	return tm.Year()
}

// Month returns the month 1-12
func (tt TinyTime) Month() time.Month {
	tm := tt.ToTime()
	return tm.Month()
}

// Day retruns the day of the month as an integer
func (tt TinyTime) Day() int {
	tm := tt.ToTime()
	return tm.Day()
}

// String returns the month formatted as a readable string
func (tt TinyTime) String() string {
	return tt.Format(time.RFC3339)
}

// IsZero returns true if the time represents the the unix epoch time
func (tt TinyTime) IsZero() bool {
	return tt.unix == 0
}

// Unix returns the time as a Unix timestamp
// where the precision is only to the day. Hours, minutes, and seconds
// are just padded zeros
func (tt TinyTime) Unix() int64 {
	t := tt.ToTime()
	return t.Unix()
}

// UnixNano returns the time as a Unix timestamp in nanoseconds
// where the precision is only to the day. Hours, minutes, seconds and nanoseconds
// are just padded zeros
func (tt TinyTime) UnixNano() int64 {
	t := tt.ToTime()
	return t.UnixNano()
}

// Weekday returns the day of the week
func (tt TinyTime) Weekday() time.Weekday {
	t := tt.ToTime()
	return t.Weekday()
}

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years, and [1,366] in leap years
func (tt TinyTime) YearDay() int {
	t := tt.ToTime()
	return t.YearDay()
}

// ISOWeek returns the ISO 8601 year and week number in which t occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (tt TinyTime) ISOWeek() (year, week int) {
	t := tt.ToTime()
	return t.ISOWeek()
}

// Date returns the readable numerical values of the date
func (tt TinyTime) Date() (year int, month time.Month, day int) {
	tm := tt.ToTime()
	return tm.Date()
}
