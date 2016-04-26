package rand

// Bern A variable following a normal distribution
type Bern struct {
	Discrete
	p float64
	// 1-p
	q float64
}

// NewBern Generate a new variable following a Gaussian law
func NewBern(p float64) (x *Bern) {
	x = &Bern{Discrete{}, p, 1. - p}
	// Discrete.discreter self reference X
	// It allows to implement more general variables
	x.Discrete = Discrete{x}
	return
}

// NewStdBern Generate a new variable following a centered, reduced Gaussian law
// (mu=0, sigma=1)
func NewStdBern() (x *Bern) {
	x = NewBern(0.5)
	return
}

// R Generate a random value following the same Bernouilli distribution as x
func (x Bern) R() int {
	var u = NewStdUnif()
	// maybe there is a more idomatic way link int(u.R()<p)
	if u.R() < x.p {
		return 1
	}
	return 0
}

// D Density probability function of the Bernouilli distribution
func (x Bern) D(k int) float64 {
	if k == 0 {
		return x.p
	} else if k == 1 {
		return x.q
	}
	return 0.
}
