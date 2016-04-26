package rand

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewUnif(t *testing.T) {
    var x = NewUnif(-1.,6.)
    assert.Equal(t, -1., x.min)
    assert.Equal(t, 6., x.max)
    assert.Equal(t, 7., x.interval)
}
func TestNewStdUnif(t *testing.T) {
    var x = NewStdUnif()
    assert.Equal(t, 0., x.min)
    assert.Equal(t, 1., x.max)
}

func TestUnifR(t *testing.T){
    var x = NewStdUnif()
    var x1 = x.R()
    assert.NotNil(t,x1)
    assert.True(t, x1>=0.)
    assert.True(t, x1<=1.)
}
func TestUnifRn(t *testing.T){
    var x = NewStdUnif()
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with min=0,max=1 : %v", x1)
}
func TestUnifD(t *testing.T){
    var x = NewStdUnif()
    var v1 = x.D(0.5)
    var v2 = x.D(1.0)
    var v3 = x.D(0.0)
    var eps = 10e-10;
    assert.InDelta(t, 1., v1, eps)
    assert.InDelta(t, 1., v2, eps)
    assert.InDelta(t, 1., v3, eps)
}
func TestUnifDOutsideInterval(t *testing.T){
    var x = NewStdUnif()
    var v1 = x.D(-1.)
    var v2 = x.D(3.)
    assert.Equal(t, 0., v1)
    assert.Equal(t, 0., v2)
}

func TestUnifP(t *testing.T){
}
func TestUnifQ(t *testing.T){
}
