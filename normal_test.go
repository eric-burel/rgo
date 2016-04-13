package rgo

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewNormal(t *testing.T) {
    var x = NewNormal(20,2)
    assert.Equal(t, 20., x.mu)
    assert.Equal(t, 2., x.sigma)
}
func TestNewStdNormal(t *testing.T) {
    var x = NewStdNormal()
    assert.Equal(t, 0., x.mu)
    assert.Equal(t, 1., x.sigma)
}

func TestNormalR(t *testing.T){
    var x = NewStdNormal()
    var x1 = x.R()
    assert.NotNil(t,x1)
}
func TestNormalRn(t *testing.T){
    var x = NewStdNormal()
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with mu = 0, sigma = 1 : %v", x1)
}
func TestNormalDSmall(t *testing.T){
    var x = NewStdNormal()
    var v1 = x.D(0.01)
    var v2 = x.D(0.005)
    var v3 = x.D(1/100000)
    var eps = 10e-8;
    // values given by R
    assert.InDelta(t, 0.3989223, v1, eps)
    assert.InDelta(t, 0.3989373, v2, eps)
    assert.InDelta(t, 0.3989423, v3, eps)
}
func TestNormalD(t *testing.T){
    var x = NewStdNormal()
    var v1 = x.D(0.)
    var v2 = x.D(0.5)
    var v3 = x.D(1.)
    var v4 = x.D(15)
    var eps = 10e-8;
    assert.InDelta(t, 0.3989423, v1, eps)
    assert.InDelta(t, 0.3520653, v2, eps)
    assert.InDelta(t, 0.2419707, v3, eps)
    assert.InDelta(t, 5.53071e-50, v4, eps)
}
func TestNormalDHuge(t *testing.T){
    var x = NewStdNormal()
    var v1 = x.D(100)
    var v2 = x.D(20000)
    var v3 = x.D(50000000000)
    var eps = 10e-8;
    assert.InDelta(t, 0., v1, eps)
    assert.InDelta(t, 0., v2, eps)
    assert.InDelta(t, 0., v3, eps)
}

func TestNormalP(t *testing.T){
}
func TestNormalQ(t *testing.T){
}
