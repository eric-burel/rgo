package rgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFiniteInt(t *testing.T) {

	var x = NewFiniteInt(map[int]float64{-1: 1. / 3., 2: 1. / 6., 3: 1. / 6., 4: 1. / 3.})
	assert.Equal(t, 1./3., x.Faces[-1])
	assert.Equal(t, 1./3., x.Faces[4])
	assert.Equal(t, 1./6., x.Faces[2])
	assert.Equal(t, 1./6., x.Faces[3])
}
func TestNewStdFiniteInt(t *testing.T) {
	var x = NewStdFiniteInt()
	var p = 1. / 6.
	assert.Equal(t, p, x.Faces[1])
	assert.Equal(t, p, x.Faces[2])
	assert.Equal(t, p, x.Faces[3])
	assert.Equal(t, p, x.Faces[4])
	assert.Equal(t, p, x.Faces[5])
	assert.Equal(t, p, x.Faces[6])
}

func TestFiniteIntR(t *testing.T) {
	var x = NewStdFiniteInt()
	var x1 = x.R()
	assert.NotNil(t, x1)
	assert.True(t, x1 >= 1)
	assert.True(t, x1 <= 6)
}
func TestFiniteIntRn(t *testing.T) {
	var x = NewStdFiniteInt()
	var x1 = x.Rn(20)
	assert.Equal(t, 20, len(x1))
	t.Logf("Generated values a standard 6 faced die : %v", x1)
}
func TestFiniteIntD(t *testing.T) {
	var x = NewFiniteInt(map[int]float64{1: 1. / 3., 2: 1. / 6., 3: 1. / 2.})
	assert.Equal(t, 1./3., x.D(1))
	assert.Equal(t, 1./6., x.D(2))
	assert.Equal(t, 1./2., x.D(3))
}
func TestFiniteIntDLim(t *testing.T) {
	var x = NewFiniteInt(map[int]float64{1: 1. / 3., 2: 1. / 6., 3: 1. / 2.})
	assert.Equal(t, 0., x.D(-1))
	assert.Equal(t, 0., x.D(4))
}
func TestFiniteIntPLim(t *testing.T) {
	var x = NewFiniteInt(map[int]float64{1: 1. / 3., 2: 1. / 6., 3: 1. / 2.})
	assert.Equal(t, 0., x.P(-1))
	assert.Equal(t, 1., x.P(4))
}

func TestFiniteIntP(t *testing.T) {
	var x = NewFiniteInt(map[int]float64{1: 1. / 3., 2: 1. / 6., 3: 1. / 2.})
	assert.Equal(t, 1./3., x.P(1))
	assert.Equal(t, 1./2., x.P(2))
	assert.Equal(t, 1., x.P(3))
}
