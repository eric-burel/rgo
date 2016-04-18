package rgo
import(
    "math/rand"
    "math"
)
// Exp A variable following a normal distribution
type Exp struct{
    Continuous
    lambda float64
}

// NewExp Generate a new variable following a Exponential law
func NewExp(lambda float64) (x *Exp){
    x = &Exp{Continuous{},lambda}
    // Discrete.discreter self reference X
    // It allows to implement more general variables
    x.Continuous = Continuous{x}
    return
}

// NewStdExp Generate a new variable following an Exponential law with rate=1
// (mu=0, sigma=1)
func NewStdExp() (x *Exp){
    x = NewExp(1.)
    return
}

// R Generate a random value following the same Exp distribution as x
func (x Exp) R() (res float64){
    res = rand.ExpFloat64() / x.lambda
    return
}

// D Density probability function of the the Exponential distribution
func (x Exp) D(v float64) (f float64){
    f = x.lambda * math.Exp(-x.lambda*v)
    return
}
