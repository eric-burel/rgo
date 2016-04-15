package rgo
import(
    "math/rand"
    "math"
)
// Normal A variable following a normal distribution
type Normal struct{
    Continuous
    mu float64
    sigma float64
}

// NewNormal Generate a new variable following a Gaussian law
func NewNormal(mu float64, sigma float64) (x *Normal){
    x = &Normal{Continuous{},mu,sigma}
    // Discrete.discreter self reference X
    // It allows to implement more general variables
    x.Continuous = Continuous{x}
    return
}

// NewStdNormal Generate a new variable following a centered, reduced Gaussian law
// (mu=0, sigma=1)
func NewStdNormal() (x *Normal){
    x = NewNormal(0.,1.)
    return
}

// R Generate a random value following the same Normal distribution as x
func (x Normal) R() (res float64){
    res = rand.NormFloat64() * x.sigma + x.mu
    return
}

// D Density probability function of the the Gaussian distribution
func (x Normal) D(v float64) (f float64){
    // TODO : calculate me once forall and multiply me with 1/x.sigma...
    norm := 1./(x.sigma*math.Sqrt2*math.SqrtPi)
    exponentiate := -0.5*math.Pow((v-x.mu)/x.sigma,2)
    f = norm*math.Exp(exponentiate)
    return
}
