package rgo

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewGeom(t *testing.T) {
    var x = NewGeom(0.6)
    assert.Equal(t, 0.6, x.p)
    assert.Equal(t, 0.4, x.q)
}
func TestNewStdGeom(t *testing.T) {
    var x = NewStdGeom()
    assert.Equal(t, 0.5, x.p)
    assert.Equal(t, 0.5, x.q)
}

func TestGeomR(t *testing.T){
    var x = NewStdGeom()
    var x1 = x.R()
    assert.NotNil(t,x1)
    assert.True(t, x1>=0)
}
func TestGeomRn(t *testing.T){
    var x = NewGeom(0.6)
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with p=0.6: %v", x1)
}
func TestGeomD(t *testing.T){
    var x = NewStdGeom()
    var eps = 10e-8;
    // values given by R
    assert.InDelta(t, 0.5, x.D(0), eps)
    assert.InDelta(t, 0.25, x.D(1), eps)
    assert.InDelta(t, 0.0625, x.D(3), eps)
    assert.InDelta(t, 0.0004882812, x.D(10), eps)
    assert.InDelta(t, 4.768372e-07, x.D(20), eps)
}
func TestGeomD06(t *testing.T){
    var x = NewGeom(0.6)
    var eps = 10e-8;
    // values given by R
    assert.InDelta(t, 0.6, x.D(0), eps)
    assert.InDelta(t, 0.096, x.D(2), eps)
    assert.InDelta(t, 6.442451e-07, x.D(15), eps)
}

func TestGeomP(t *testing.T){
    var x = NewStdGeom()
    var eps = 10e-8;
    // values given by R
    assert.InDelta(t, 0.5, x.P(0), eps)
    assert.InDelta(t, 0.75, x.P(1), eps)
    assert.InDelta(t,  0.984375, x.P(5), eps)
    assert.InDelta(t, 0.9995117, x.P(10), eps)
}

func TestGeomLims(t *testing.T){
    var x = NewStdGeom()
    assert.Equal(t,0., x.D(-1))
    assert.Equal(t,0., x.P(-1))
}
func TestGeomQ(t *testing.T){
}
