package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatImplementSampler(t *testing.T) {
	//@see https://splice.com/blog/golang-verify-type-implements-interface-compile-time/
	// this won't compile if Int does not implement Sampler
	var _ Sampler = (*Float)(nil)
}
func TestFloatImplementInterface(t *testing.T) {
	var _ Interface = (*Float)(nil)
}
func TestFloatImplementFloater(t *testing.T) {
	var _ Floater = (*Float)(nil)
}
func TestFloatNewSampleFloat(t *testing.T) {
	x := NewFloat([]float64{1., 2., 3.})
	assert.NotNil(t, x)
}

func TestFloatSum(t *testing.T) {
	x := NewFloat([]float64{1., 2., 3.})
	sum := x.Sum()
	assert.Equal(t, 6., sum)
}

func TestFloatProd(t *testing.T) {
	x := NewFloat([]float64{2., 2., 3.})
	prod := x.Prod()
	assert.Equal(t, 12., prod)
}

func TestFloatLength(t *testing.T) {
	x := NewFloat([]float64{2., 2., 3.})
	assert.Equal(t, 3, x.Length())
}

func TestFloatMin(t *testing.T) {
	x := NewFloat([]float64{2., -2., 3.})
	assert.Equal(t, -2., x.Min())
}
func TestFloatMax(t *testing.T) {
	x := NewFloat([]float64{2., -2., 3.})
	assert.Equal(t, 3., x.Max())
}

func TestFloatArgmin(t *testing.T) {
	x := NewFloat([]float64{2., -2., 3.})
	assert.Equal(t, 1, x.Argmin())
}
func TestFloatArgminAll(t *testing.T) {
	x := NewFloat([]float64{-2., -2., 3.})
	assert.Equal(t, []int{0., 1.}, x.ArgminAll())
}
func TestFloatArgmax(t *testing.T) {
	x := NewFloat([]float64{-2., 4., 4.})
	assert.Equal(t, 1, x.Argmax())
}
func TestFloatArgmaxAll(t *testing.T) {
	x := NewFloat([]float64{4., 4., 3.})
	assert.Equal(t, []int{0, 1}, x.ArgmaxAll())
}
func TestFloatMedian(t *testing.T) {
	x1 := NewFloat([]float64{1., 2., 3.})
	assert.Equal(t, 2., x1.Median())
	x2 := NewFloat([]float64{1., 2.})
	assert.Equal(t, 1.5, x2.Median())
}

func TestFloatVar(t *testing.T) {
	x := NewFloat([]float64{4., 2., 0.})
	// sqrt 4 + 4
	assert.InEpsilon(t, 8./2., x.Var(), 10e-8)
}

/*func TestFloatCov(t *testing.T) {
	x := NewFloat([]float64{1, 2, 3})
	y := NewFloat([]float64{6, 5, 37})
	assert.Equal(t, 15.5, x.Cov(*y))
}*/

func TestFloatEmptySample(t *testing.T) {
	x := NewFloat([]float64{})
	assert.Equal(t, 0., x.Sum())
	assert.Equal(t, 1., x.Prod())
	assert.Panics(t, func() { x.Mean() })
	assert.Panics(t, func() { x.Var() })
	assert.Panics(t, func() { x.Min() })
	assert.Panics(t, func() { x.Max() })
	assert.Panics(t, func() { x.Argmin() })
	assert.Panics(t, func() { x.Argmax() })
	assert.Panics(t, func() { x.ArgminAll() })
	assert.Panics(t, func() { x.ArgmaxAll() })
}
func TestFloatSizeOne(t *testing.T) {
	x := NewFloat([]float64{2.})
	assert.Equal(t, 2., x.Sum())
	assert.Equal(t, 2., x.Prod())
	assert.Equal(t, 2., x.Mean()) // is that a correct math definition ?
	assert.Equal(t, 0., x.Var())
	assert.Equal(t, 2., x.Min())
	assert.Equal(t, 2., x.Max())
	assert.Equal(t, 0, x.Argmin())
	assert.Equal(t, 0, x.Argmax())
	assert.Equal(t, []int{0}, x.ArgminAll())
	assert.Equal(t, []int{0}, x.ArgmaxAll())
}
