package tinytime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalBinary(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	dat, err := tt.MarshalBinary()
	assert.Nil(t, err)
	assert.Equal(t, []byte{
		tinyTimeBinaryVersion,
		byte(94),
		byte(132),
		byte(161),
		byte(102),
	}, dat)

	newtt := TinyTime{}
	err = newtt.UnmarshalBinary(dat)
	assert.Nil(t, err)
	assert.Equal(t, tt, newtt)
}

func TestUnmarshalBinary(t *testing.T) {
	dat := []byte{
		tinyTimeBinaryVersion,
		byte(94),
		byte(132),
		byte(161),
		byte(102),
	}
	tt := TinyTime{}
	err := tt.UnmarshalBinary(dat)
	assert.Nil(t, err)
	assert.Equal(t, TinyTime{unix: 1585750374}, tt)

	newDat, _ := tt.MarshalBinary()
	assert.Equal(t, dat, newDat)
}
