package tinytime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromTimeValid(t *testing.T) {
	tm := time.Date(2020, time.Month(4), 1, 14, 12, 54, 0, time.UTC)
	tt, err := FromTime(tm)
	assert.Nil(t, err)
	assert.Equal(t, TinyTime{unix: 1585750374}, tt)
}

func TestFromTimeInvalid(t *testing.T) {
	tm := time.Date(100000, time.Month(5), 2, 0, 0, 0, 0, time.UTC)
	_, err := FromTime(tm)
	assert.NotNil(t, err)
}

func TestToTime(t *testing.T) {
	tt := TinyTime{unix: 1585754374}
	tm := tt.ToTime()
	assert.Equal(t, time.Date(2020, 4, 1, 15, 19, 34, 0, time.UTC), tm)

	tt = TinyTime{unix: 1585750374}
	tm = tt.ToTime()
	assert.Equal(t, time.Date(2020, 4, 1, 14, 12, 54, 0, time.UTC), tm)
}
