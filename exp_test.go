package rgo

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewExp(t *testing.T) {
    var x = NewExp(10)
    assert.Equal(t, 10., x.lambda)
}
func TestNewStdExp(t *testing.T) {
    var x = NewStdExp()
    assert.Equal(t, 1., x.lambda)
}

func TestExpR(t *testing.T){
    var x = NewStdExp()
    var x1 = x.R()
    assert.NotNil(t,x1)
}
func TestExpRn(t *testing.T){
    var x = NewStdExp()
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with lambda = 1. : %v", x1)
}
func TestExpD(t *testing.T){
    var x = NewStdExp()
    var eps = 10e-8;
    // values given by R
    assert.InDelta(t, 0.9512294, x.D(0.05), eps)
    assert.InDelta(t, 0.9048374, x.D(0.1), eps)
    assert.InDelta(t, 0.6065307, x.D(0.5), eps)
    assert.InDelta(t, 0.3678794, x.D(1), eps)
    assert.InDelta(t, 0.2231302, x.D(1.5), eps)
    assert.InDelta(t, 0.1353353, x.D(2), eps)
    assert.InDelta(t, 4.539993e-05, x.D(10), eps)
    assert.InDelta(t, 9.357623e-14, x.D(30), eps)
}
func TestExpDSmall(t *testing.T){
    var x = NewStdExp()
    var eps = 10e-8;
    assert.InDelta(t, 0.9990005, x.D(0.001), eps)
}
func TestExpD10(t *testing.T){
    var x = NewExp(10)
    var eps = 10e-8;
    assert.InDelta(t, 0.06737947, x.D(0.5), eps)
    assert.InDelta(t, 0.0004539993, x.D(1), eps)
    assert.InDelta(t,  3.720076e-43, x.D(10), eps)
}

func TestExpP(t *testing.T){
}
func TestExpQ(t *testing.T){
}
