package sample

// Float A sample of float64 values
type Float []float64

// Floater Functions implemented by a sample of float64
// NOTE : not that useful...
type Floater interface {
	Sum() float64
	Prod() float64
	Max() float64
	Min() float64
}

// NewFloat Generate a float sample given an array of float64
func NewFloat(x []float64) (s *Float) {
	s = new(Float)
	*s = Float(x)
	return
}

// implement sort
func (s Float) Len() int {
	return len(s)
}

func (s Float) Less(i, j int) bool {
	return s[i] <= s[j]
}
func (s Float) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
	return
}

// GetValueFloat64 Retrieve ith element and converts it to a float64
func (s Float) GetValueFloat64(index int) (value float64) {
	value = float64(s[index])
	return
}

// Sum Calculates the sum of the elements
func (s Float) Sum() (sum float64) {
	return SumFloat64(s)
}

// Prod Calculates the products of each element
func (s Float) Prod() (prod float64) {
	return ProdFloat64(s)
}

func (s Float) Length() (length int) {
	length = len(s)
	return
}

func (s Float) Mean() (mean float64) {
	mean = Mean(s)
	return
}

func (s Float) Min() (min float64) {
	return MinFloat64(s)
}
func (s Float) Max() (max float64) {
	return MaxFloat64(s)
}

func (s Float) Argmin() (argmin int) {
	return Argmin(s)
}

// Argmax Return a single argmax value (first match)
func (s Float) Argmax() (argmax int) {
	return Argmax(s)
}
func (s Float) ArgminAll() (argmin []int) {
	return ArgminAll(s)
}
func (s Float) ArgmaxAll() (argmax []int) {
	return ArgmaxAll(s)
}

func (s Float) Median() (mu float64) {
	return Median(s)
}
func (s Float) Quantile(p float64) (mu float64) {
	return Quantile(s, p)
}

// Var Unbiased variance calculation
// NOTE : Beware ! this is NOT sqrt(cov(x,x)), as it is divided by n-1 not by n
func (s Float) Var() (sigma2 float64) {
	return Var(s)
}
