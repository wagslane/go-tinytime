package tinytime

import (
	"fmt"
	"time"
)

// Parse a layout into a TinyTime
func Parse(layout, value string) (TinyTime, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return TinyTime{}, fmt.Errorf("tinytime Parse: %v", err)
	}
	tt, err := FromTime(t)
	if err != nil {
		return TinyTime{}, fmt.Errorf("tinytime Parse: %v", err)
	}
	return tt, nil
}

// ParseInLocation parses a layout into a TinyTime including a location.
// The input is converted into UTC
func ParseInLocation(layout, value string, loc *time.Location) (TinyTime, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return TinyTime{}, fmt.Errorf("tinytime ParseInLocation: %v", err)
	}
	tt, err := FromTime(t)
	if err != nil {
		return TinyTime{}, fmt.Errorf("tinytime ParseInLocation: %v", err)
	}
	return tt, nil
}

// AppendFormat is like Format but appends the textual representation to b and returns the extended buffer
func (tt TinyTime) AppendFormat(b []byte, layout string) []byte {
	t := tt.ToTime()
	return t.AppendFormat(b, layout)
}

// Format returns a formatted time, as specified by the standard time library
// https://golang.org/src/time/format.go?s=16029:16071#L485
func (tt TinyTime) Format(layout string) string {
	t := tt.ToTime()
	return t.Format(layout)
}
