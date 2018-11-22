package gonn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	// get the coeffient of the plane
	b, err := GetBoundary([]float32{1.0, 0.0}, []float32{0.0, 1.0})
	assert.Equal(t, err, nil)
	// evaluate the value of (1.0, 1.0) on the plane
	v, err := EvalFormula(b, []float32{1.0, 1.0})
	assert.Equal(t, err, nil)
	assert.Equal(t, v, float32(0.0))
}
