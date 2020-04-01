package tinytime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tt, err := Parse(time.RFC3339, "2020-04-01T14:12:54Z")
	assert.Nil(t, err)
	assert.Equal(t, 1585750374, int(tt.unix))
}

func TestParseInLocation(t *testing.T) {
	tt, err := ParseInLocation(time.RFC3339, "2020-04-01T14:12:54Z", time.UTC)
	assert.Nil(t, err)
	assert.Equal(t, 1585750374, int(tt.unix))
}

func TestAppendFormat(t *testing.T) {
	tt := New(1585750374)
	text := []byte("Time: ")

	text = tt.AppendFormat(text, time.RFC3339)
	assert.Equal(t, "Time: 2020-04-01T14:12:54Z", string(text))
}

func TestFormat(t *testing.T) {
	tt := New(1585750374)

	assert.Equal(t, "2020-04-01T14:12:54Z", string(tt.Format(time.RFC3339)))
}
