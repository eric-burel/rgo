package rand

// FiniteInt A variable following a binomial distribution
type FiniteInt struct {
	Discrete
	Faces map[int]float64
}

// NewFiniteInt Generate a new variable following a binomial law
func NewFiniteInt(faces map[int]float64) (x *FiniteInt) {
	x = &FiniteInt{Discrete{}, faces}
	x.Discrete = Discrete{x}
	return
}

// NewStdFiniteInt Generate a new six faces dice
func NewStdFiniteInt() (x *FiniteInt) {
	p := 1. / 6.
	x = NewFiniteInt(map[int]float64{1: p, 2: p, 3: p, 4: p, 5: p, 6: p})
	return
}

// R Generate a random value following the same FiniteInt distribution as x
// NOTE res is always defined, even if all probabilities does not add to 1
// in this case the last value is sent, making it more probable
// (eventhough it won't be noticeable in an usual context)
func (x FiniteInt) R() (res int) {
	// Inversion, might be a bit costly
	u := NewStdUnif().R()
	sumP := 0.
	for value, prob := range x.Faces {
		res = value
		sumP += prob
		if sumP > u {
			break
		}
	}
	return
}

// D Density probability function of the FiniteInt distribution P(X=k)=p*(1-p)^k
func (x FiniteInt) D(k int) (res float64) {
	res = 0.
	for value, prob := range x.Faces {
		if k == value {
			res += prob
		}
	}
	return
}

// P Cumulative density function of the FiniteInt distribution
// If probabilities does not really add up, it will still return 1 if k is greater
// or equal to all values defined in the law
// Ie for a standard die, D(6) = 1.0 even if 1./6 + ... + 1./6. does not really equal 1.0
func (x FiniteInt) P(k int) (res float64) {
	res = 0.
	all := true
	for value, prob := range x.Faces {
		if k >= value {
			res += prob
		} else {
			all = false
		}
	}
	// In case the probabilities do not add up
	// we check if k was over all the values, meaning that k is outside the variable support intervael
	if all {
		res = 1.
	}
	return
}
