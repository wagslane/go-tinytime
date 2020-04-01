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

// Clock returns the hour, minute, and second within the day specified by tt
func (tt TinyTime) Clock() (hour, min, sec int) {
	tm := tt.ToTime()
	return tm.Clock()
}

// Hour returns the hour within the day specified by tt, in the range [0, 23]
func (tt TinyTime) Hour() int {
	tm := tt.ToTime()
	return tm.Hour()
}

// Minute returns the minute offset within the hour specified by tt, in the range [0, 59]
func (tt TinyTime) Minute() int {
	tm := tt.ToTime()
	return tm.Minute()
}

// Second returns the second offset within the minute specified by tt, in the range [0, 59]
func (tt TinyTime) Second() int {
	tm := tt.ToTime()
	return tm.Second()
}

// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (tt TinyTime) Round(d time.Duration) time.Time {
	tm := tt.ToTime()
	return tm.Round(d)
}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (tt TinyTime) Truncate(d time.Duration) time.Time {
	tm := tt.ToTime()
	return tm.Truncate(d)
}
