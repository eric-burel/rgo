package rgo


import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsProba(t *testing.T){
    assert.True(t, IsProba(0.5))
    assert.True(t, IsProba(0.0))
    assert.True(t, IsProba(1.0))

    assert.False(t, IsProba(-5))
    assert.False(t, IsProba(1.1))
    assert.False(t, IsProba(-0.1))
}

func TestSumInt64(t * testing.T){
    var array = IntArr{10,10,20};
    assert.Equal(t, int(40), array.Sum())
}
func TestSumFloat64(t * testing.T){
    var array = Float64Arr{10.,10.,20.}
    assert.Equal(t, 40., array.Sum())
}
