package rgo

// Binom A variable following a binomial distribution
type Binom struct{
    Discrete
    n int
    p float64
    // 1-p
    q float64
}


// NewBinom Generate a new variable following a binomial law
func NewBinom(n int, p float64) (x *Binom){
    x = &Binom{Discrete{},n,p,1.-p}
    x.Discrete = Discrete{x};
    return
}

// NewStdBinom Generate a new variable following a B(1,0.5) law (equal to Bern(0.5))
// (mu=0, sigma=1)
func NewStdBinom() (x *Binom){
    x = NewBinom(1,0.5)
    return
}

// R Generate a random value following the same Binomouilli distribution as x
func (x Binom) R() (res int) {
    var b = NewBern(x.p)
    // sums the result of n Bern(p) draws and sums them
    res = intArr(b.Rn(int(x.n))).Sum()
    return
}

// D Density probability function of the Binom distribution
func (x Binom) D(k int) float64{
    // TODO
    return 0.
}
