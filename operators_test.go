package tinytime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAfter(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	tu := TinyTime{unix: 1585750374}
	assert.False(t, tt.After(tu))

	tt = TinyTime{unix: 1585750375}
	tu = TinyTime{unix: 1585750376}
	assert.False(t, tt.After(tu))

	tt = TinyTime{unix: 1585750374}
	tu = TinyTime{unix: 1585750373}
	assert.True(t, tt.After(tu))
}

func TestBefore(t *testing.T) {
	tt := TinyTime{unix: 1585750374}
	tu := TinyTime{unix: 1585750374}
	assert.False(t, tt.Before(tu))

	tt = TinyTime{unix: 1585750374}
	tu = TinyTime{unix: 1585750373}
	assert.False(t, tt.Before(tu))

	tt = TinyTime{unix: 1585750374}
	tu = TinyTime{unix: 1585750375}
	assert.True(t, tt.Before(tu))
}

func TestAdd(t *testing.T) {
	tt := TinyTime{unix: 1583020800}
	assert.Equal(t,
		TinyTime{unix: 1583107200},
		tt.Add(time.Hour*24),
	)

	tt = TinyTime{unix: 1583107200}
	assert.Equal(t,
		TinyTime{unix: 1583020800},
		tt.Add(-time.Hour*24),
	)
}

func TestAddDate(t *testing.T) {
	tt := TinyTime{unix: 1583020800}
	assert.Equal(t,
		TinyTime{unix: 1583107200},
		tt.AddDate(0, 0, 1),
	)

	tt = TinyTime{unix: 1583107200}
	assert.Equal(t,
		TinyTime{unix: 1583020800},
		tt.AddDate(0, 0, -1),
	)
}

func TestEqual(t *testing.T) {
	tt := New(1585750374)
	assert.True(t, tt.Equal(tt))

	tt = New(1583107200)
	tu := New(1583020800)
	assert.False(t, tt.Equal(tu))
}

func TestSub(t *testing.T) {
	tt := New(1583107200)
	tu := New(1583020800)
	assert.Equal(t, time.Hour*24, tt.Sub(tu))

	tt = New(1585750374)
	tu = New(1585750374)
	assert.Equal(t, time.Duration(0), tt.Sub(tu))

	tt = New(1583020800)
	tu = New(1583107200)
	assert.Equal(t, -time.Hour*24, tt.Sub(tu))
}
