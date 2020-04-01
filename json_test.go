package tinytime

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	dat, err := TinyTime{unix: 1585750374}.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, `"2020-04-01T14:12:54Z"`, string(dat))

	s := struct {
		TD TinyTime
	}{
		TD: TinyTime{unix: 1585750374},
	}
	dat, err = json.Marshal(s)
	assert.Nil(t, err)
	assert.Equal(t, `{"TD":"2020-04-01T14:12:54Z"}`, string(dat))
}

func TestUnmarshalJSON(t *testing.T) {
	dat := []byte(`"2020-04-01T14:12:54Z"`)
	tt := TinyTime{}
	err := tt.UnmarshalJSON(dat)
	assert.Nil(t, err)
	assert.Equal(t, TinyTime{unix: 1585750374}, tt)

	dat = []byte(`{"TD":"2020-04-01T14:12:54Z"}`)
	s := struct {
		TD TinyTime
	}{}
	err = json.Unmarshal(dat, &s)
	assert.Nil(t, err)
	assert.Equal(t, struct {
		TD TinyTime
	}{
		TD: TinyTime{unix: 1585750374},
	}, s)
}
