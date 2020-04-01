package tinytime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalText(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	dat, err := tt.MarshalText()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`2020-04-01T14:12:54Z`), dat)
}

func TestUnmarshalText(t *testing.T) {
	dat := []byte(`2020-04-01T14:12:54Z`)
	tt := TinyTime{}
	err := tt.UnmarshalText(dat)
	assert.Nil(t, err)
	assert.Equal(t, TinyTime{unix: 1585750374}, tt)
}
