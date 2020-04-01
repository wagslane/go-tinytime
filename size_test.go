package tinytime

import (
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSizes(t *testing.T) {
	tt := TinyTime{}
	assert.Equal(t, uintptr(4), unsafe.Sizeof(tt))

	tm := time.Now().UTC()
	assert.Equal(t, uintptr(24), unsafe.Sizeof(tm))
}
