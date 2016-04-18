
package rgo

import(
    "math"
)
// Geom A variable following a binomial distribution
type Geom struct{
    Discrete
    p float64
    q float64
}


// NewGeom Generate a new variable following a binomial law
func NewGeom(p float64) (x *Geom){
    x = &Geom{Discrete{},p,1.-p}
    x.Discrete = Discrete{x};
    return
}

// NewStdGeom Generate a new variable following a B(1,0.5) law (equal to Bern(0.5))
// (mu=0, sigma=1)
func NewStdGeom() (x *Geom){
    x = NewGeom(0.5)
    return
}

// R Generate a random value following the same Geomouilli distribution as x
func (x Geom) R() (res int) {
    // Simulation by inversion, as F-1 is easy to calculate in this case
    var u = NewStdUnif()
    return x.invP(u.R())
}

// D Density probability function of the Geom distribution P(X=k)=p*(1-p)^k
// NOTE : One might prefer P(X=0) = 0, and P(X=k)=p*(1-p)^(k-1)
// However we chose the same definition as R
func (x Geom) D(k int) float64{
    if (k < 0){
        return 0
    }
    return math.Pow(x.q, float64(k) )*x.p
}

// P Cumulative density function of the Geom distribution
func (x Geom) P(k int) float64{
    if (k < 0){
        return 0
    }
    return 1. - math.Pow(x.q, float64(k+1))
}

// Inverse of the cumulative density function
func (x Geom) invP(p float64) int{
    return int(math.Log10(1.-p)/math.Log(x.q))
}
