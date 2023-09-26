package tinytime

import (
	"errors"
	"math"
	"time"
)

// FromTime converts a time.Time into a TinyTime
// all tinyTinyTime.TinyTimes are UTC timezone, so that conversion will
// be made here
func FromTime(t time.Time) (TinyTime, error) {
	t = t.UTC()
	unix := t.Unix()
	if unix > math.MaxUint32 {
		return TinyTime{}, errors.New("tinytime FromTime: unix must be less than math.MaxUint32")
	}
	return TinyTime{
		unix: uint32(unix),
	}, nil
}

// ToTime converts a tinyTinyTime.TinyTime to a time.Time, always UTC
func (tt TinyTime) ToTime() time.Time {
	return time.Unix(int64(tt.unix), 0).UTC()
}
