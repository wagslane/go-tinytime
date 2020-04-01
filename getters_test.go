package tinytime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	assert.Equal(t, 2020, tt.Year())
}

func TestMonth(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	assert.Equal(t, time.April, tt.Month())
}

func TestDay(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	assert.Equal(t, 1, tt.Day())
}

func TestString(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	assert.Equal(t, "2020-04-01T14:12:54Z", tt.String())
}

func TestIsZero(t *testing.T) {
	assert.True(t, TinyTime{}.IsZero())
	assert.True(t, TinyTime{unix: 0}.IsZero())
	assert.False(t, TinyTime{unix: 1}.IsZero())
}

func TestToUnix(t *testing.T) {
	tt := New(1585750374)
	assert.Equal(t, int64(1585750374), tt.Unix())

	tt = New(0)
	assert.Equal(t, int64(0), tt.Unix())
}

func TestToUnixNano(t *testing.T) {
	tt := New(1585750374)
	assert.Equal(t, int64(1585750374000000000), tt.UnixNano())

	tt = New(0)
	assert.Equal(t, int64(0), tt.UnixNano())
}

func TestWeekday(t *testing.T) {
	tt := New(1585750374)
	assert.Equal(t, time.Wednesday, tt.Weekday())
}

func TestYearDay(t *testing.T) {
	tt := New(1577836800)
	assert.Equal(t, 1, tt.YearDay())

	tt = New(1585750374)
	assert.Equal(t, 92, tt.YearDay())
}

func TestISOWeek(t *testing.T) {
	tt := New(1585750374)
	year, week := tt.ISOWeek()
	assert.Equal(t, 2020, year)
	assert.Equal(t, 14, week)
}

func TestDate(t *testing.T) {
	tt := New(1585750374)
	year, month, day := tt.Date()
	assert.Equal(t, 2020, year)
	assert.Equal(t, time.April, month)
	assert.Equal(t, 1, day)
}
