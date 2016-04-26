package rand

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewBern(t *testing.T) {
    var x = NewBern(0.6)
    assert.Equal(t, 0.6, x.p)
}
func TestNewStdBern(t *testing.T) {
    var x = NewStdBern()
    assert.Equal(t, 0.5, x.p)
}

func TestBernR(t *testing.T){
    var x = NewStdBern()
    var x1 = x.R()
    assert.NotNil(t,x1)
    assert.True(t, x1>=0)
    assert.True(t, x1<=1)
}
func TestBernRn(t *testing.T){
    var x = NewStdBern()
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with p=0.5 : %v", x1)
}
func TestBernD(t *testing.T){
    var x = NewBern(0.4)
    var v1 = x.D(0)
    var v2 = x.D(1)
    assert.Equal(t, x.p, v1)
    assert.Equal(t, x.q, v2)
}
func TestBernDOutsideInterval(t *testing.T){
    var x = NewStdBern()
    var v1 = x.D(-1)
    var v2 = x.D(3)
    assert.Equal(t, 0., v1)
    assert.Equal(t, 0., v2)
}

func TestBernP(t *testing.T){
}
func TestBernQ(t *testing.T){
}
