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
}
func TestNormalD(t *testing.T){
}
func TestNormalP(t *testing.T){
}
func TestNormalQ(t *testing.T){
}
