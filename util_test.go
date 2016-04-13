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
