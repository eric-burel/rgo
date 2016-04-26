package rand
import(
    "math/rand"
    "math"
)
// Norm A variable following a normal distribution
type Norm struct{
    Continuous
    mu float64
    sigma float64
}

// NewNorm Generate a new variable following a Gaussian law
func NewNorm(mu float64, sigma float64) (x *Norm){
    x = &Norm{Continuous{},mu,sigma}
    // Discrete.discreter self reference X
    // It allows to implement more general variables
    x.Continuous = Continuous{x}
    return
}

// NewStdNorm Generate a new variable following a centered, reduced Gaussian law
// (mu=0, sigma=1)
func NewStdNorm() (x *Norm){
    x = NewNorm(0.,1.)
    return
}

// R Generate a random value following the same Norm distribution as x
func (x Norm) R() (res float64){
    res = rand.NormFloat64() * x.sigma + x.mu
    return
}

// D Density probability function of the the Gaussian distribution
func (x Norm) D(v float64) (f float64){
    // TODO : calculate me once forall and multiply me with 1/x.sigma...
    norm := 1./(x.sigma*math.Sqrt2*math.SqrtPi)
    exponentiate := -0.5*math.Pow((v-x.mu)/x.sigma,2)
    f = norm*math.Exp(exponentiate)
    return
}
