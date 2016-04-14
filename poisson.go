package rgo

import(
    "math"
)

// Poisson A variable following a normal distribution
type Poisson struct{
    lambda float64
}

// NewPoisson Generate a new variable following a Gaussian law
func NewPoisson(lambda float64) (x *Poisson){
    x = &Poisson{lambda}
    return
}

// NewStdPoisson Generate a new variable following a centered, reduced Gaussian law
// (mu=0, sigma=1)
func NewStdPoisson() (x *Poisson){
    x = &Poisson{1.}
    return
}

// R Generate a random value following the same Poisson distribution as x
func (x Poisson) R() (res float64){
    // TODO : here is a HUGE approximation
    // plus it does work bad (results get underestimated) if lambda is under 10e-3
    // a complete process to simulate Poisson is harder to implement but possible,
    // and as been done
    var binom = NewBinom(math.Ceil(x.lambda*10e3), 10e-3)
    res = binom.R()
    return
}
// Rn Generate an array of random values following the same Poisson distribution as x
func (x Poisson) Rn(n int) (res []float64){
    res = make([]float64, n)
    for i := 0; i < n; i++ {
        res[i] = x.R()
    }
    return
}

// D Density probability function of the the Poisson distribution
func (x Poisson) D(k int64) (f float64){
    norm := math.Exp(-x.lambda)
    // this syntax avoid explosion, we multiply result by lambda, then divide it,
    // so that mult is neither too small nor too big after each step
    mult := x.lambda
    for i := 1; i <= int(k) ; i++{
        mult *= x.lambda
        mult /= float64(i)
    }
    f = norm*mult
    return
}
