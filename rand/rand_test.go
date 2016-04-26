package rand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verify that each variable use a different seed
func TestSeed(t *testing.T) {
	x := NewStdNorm()
	y := NewStdNorm()
	assert.NotEqual(t, x.R(), y.R())
}

func TestSeedDifferentRun(t *testing.T) {
	assert.Equal(t, seeded, true)
}
