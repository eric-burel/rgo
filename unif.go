package rgo
import(
    "math/rand"
)
// Unif A variable following a normal distribution
type Unif struct{
    min float64
    max float64
    interval float64
}


// NewUnif Generate a new variable following a Gaussian law
func NewUnif(min float64, max float64) (x *Unif){
    x = &Unif{min, max, max-min}
    return
}

// NewStdUnif Generate a new variable following a centered, reduced Gaussian law
// (mu=0, sigma=1)
func NewStdUnif() (x *Unif){
    x = NewUnif(0.,1.)
    return
}

// R Generate a random value following the same Unif distribution as x
func (x Unif) R() (res float64){
    res = rand.Float64() * x.interval + x.min
    return
}
// Rn Generate an array of random values following the same Unif distribution as x
func (x Unif) Rn(n int) (res []float64){
    res = make([]float64, n)
    for i := 0; i < n; i++ {
        res[i] = x.R()
    }
    return
}

// D Density probability function of the Uniform distribution
func (x Unif) D(v float64) float64{
    if v < x.min || v > x.max {
        return 0.
    }
    return 1/x.interval
}
