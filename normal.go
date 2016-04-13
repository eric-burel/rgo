package rgo

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
