package rgo
import(
    "math/rand"
)
// Normal A variable following a normal distribution
type Normal struct{
    mu float64
    sigma float64
}

// NewNormal Generate a new variable following a Gaussian law
func NewNormal(mu float64, sigma float64) (x *Normal){
    x = &Normal{mu,sigma}
    return
}

// NewStdNormal Generate a new variable following a centered, reduced Gaussian law
// (mu=0, sigma=1)
func NewStdNormal() (x *Normal){
    x = &Normal{0,1}
    return
}

// R Generate a random value following the same Normal distribution as x
func (x Normal) R() (res float64){
    res = rand.NormFloat64() * x.sigma + x.mu
    return
}
// Rn Generate an array of random values following the same Normal distribution as x
func (x Normal) Rn(n int) (res []float64){
    res = make([]float64, n)
    for i := 0; i < n; i++ {
        res[i] = rand.NormFloat64() * x.sigma + x.mu
    }
    return
}
