package rgo

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewPoisson(t *testing.T) {
    var x = NewPoisson(6)
    assert.Equal(t, 6., x.lambda)
}
func TestNewStdPoisson(t *testing.T) {
    var x = NewStdPoisson()
    assert.Equal(t, 1., x.lambda)
}

func TestPoissonR(t *testing.T){
    var x = NewStdPoisson()
    var x1 = x.R()
    assert.NotNil(t,x1)
    assert.True(t, x1>=0)
}
func TestPoissonRn(t *testing.T){
    var x = NewStdPoisson()
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with lambda=1 : %v", x1)
}
func TestPoissonD(t *testing.T){
    var x = NewStdPoisson()
    var v1 = x.D(1)
    var v2 = x.D(5)
    var v3 = x.D(10)
    var eps = 10e-8;
    assert.InDelta(t, 0.3678794, v1, eps)
    assert.InDelta(t, 0.006737947, v2, eps)
    assert.InDelta(t, 4.539993e-05, v3, eps)
}
func TestPoissonDHuge(t *testing.T){
    var x = NewStdPoisson()
    var v1 = x.D(100)
    var v2 = x.D(20000)
    var v3 = x.D(50000000000)
    var eps = 10e-8;
    assert.InDelta(t, 3.720076e-44, v1, eps)
    assert.InDelta(t, 0., v2, eps)
    assert.InDelta(t, 0., v3, eps)
}

func TestPoissonP(t *testing.T){
}
func TestPoissonQ(t *testing.T){
}
