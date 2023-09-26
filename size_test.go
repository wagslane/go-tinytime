package tinytime

import (
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSizes(t *testing.T) {
	assert.Equal(t, uintptr(4), unsafe.Sizeof(TinyTime{}))
	assert.Equal(t, uintptr(24), unsafe.Sizeof(time.Now().UTC()))
}
