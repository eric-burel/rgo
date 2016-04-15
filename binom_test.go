package rgo

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewBinom(t *testing.T) {
    var x = NewBinom(10,0.6)
    assert.Equal(t, 0.6, x.p)
    assert.Equal(t, int(10), x.n)
}
func TestNewStdBinom(t *testing.T) {
    var x = NewStdBinom()
    assert.Equal(t, 0.5, x.p)
    assert.Equal(t, 1, x.n)
}

func TestBinomR(t *testing.T){
    var x = NewStdBinom()
    var x1 = x.R()
    assert.NotNil(t,x1)
    assert.True(t, x1>=0)
    assert.True(t, x1<=1)
}
func TestBinomRn(t *testing.T){
    var x = NewBinom(10,0.5)
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with n=10 p=0.5 : %v", x1)
}
func TestBinomD(t *testing.T){
    /*var x = NewBinom(0.4)
    var v1 = x.D(0)
    var v2 = x.D(1)
    assert.Equal(t, x.p, v1)
    assert.Equal(t, x.q, v2)*/
}

func TestBinomP(t *testing.T){
}
func TestBinomQ(t *testing.T){
}
