package rand

import(
    "math"
)

// Pois A variable following a normal distribution
type Pois struct{
    lambda float64
}

// NewPois Generate a new variable following a Gaussian law
func NewPois(lambda float64) (x *Pois){
    x = &Pois{lambda}
    return
}

// NewStdPois Generate a new variable following a centered, reduced Gaussian law
// (mu=0, sigma=1)
func NewStdPois() (x *Pois){
    x = &Pois{1.}
    return
}

// R Generate a random value following the same Pois distribution as x
func (x Pois) R() (res int){
    var lambda = x.lambda;
    var power = math.Log10(lambda);
    var scale = 10e-3
    // if lambda < 0.1, we do some scaling so that lambda/scale is an integer over 1
    if (power < 0){
        scale = scale*math.Pow(10, power)
    }
    // TODO : here is a HUGE approximation
    // plus it does work bad (results get underestimated) if lambda is under 10e-3
    // a complete process to simulate Pois is harder to implement but possible,
    // and as been don  e
    var binom = NewBinom(int(math.Ceil(lambda/scale)), scale)
    res = binom.R()
    return
}
// Rn Generate an array of random values following the same Pois distribution as x
func (x Pois) Rn(n int) (res []int){
    res = make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = x.R()
    }
    return
}

// D Density probability function of the the Pois distribution
func (x Pois) D(k int) (f float64){
    norm := math.Exp(-x.lambda)
    // this syntax avoid explosion, we multiply result by lambda, then divide it,
    // so that mult is neither too small nor too big after each step
    mult := 1. // lambda^0/0!
    for i := 1; i <= int(k) ; i++{
        mult *= x.lambda
        mult /= float64(i)
    }
    f = norm*mult
    return
}
