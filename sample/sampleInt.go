package sample

// Int A sample of float values
type Int []int

// Inter Interface for Int samples
type Inter interface {
	Sum() int
	Prod() int
	Max() int
	Min() int
}

// NewInt Generate a sample of float values that can be analysed, from an array
func NewInt(x []int) (s *Int) {
	s = new(Int)
	*s = Int(x)
	return
}

// Implement sort
func (s Int) Len() int {
	return len(s)
}

// Less Compare two element of the array
func (s Int) Less(i, j int) bool {
	return s[i] <= s[j]
}

// Swap Swap two element of the array
func (s Int) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
	return
}

// GetValueInt64 Retrieve ith element and converts it to a float64
func (s Int) GetValueFloat64(index int) (value float64) {
	value = float64(s[index])
	return
}

// Sum Sum of the elements, as an int
func (s Int) Sum() (sum int) {
	return int(SumFloat64(s))
}

func (s Int) Prod() (prod int) {
	return int(ProdFloat64(s))
}

func (s Int) Length() (length int) {
	length = len(s)
	return
}

func (s Int) Mean() (mean float64) {
	mean = Mean(s)
	return
}

func (s Int) Min() (min int) {
	return int(MinFloat64(s))
}
func (s Int) Max() (max int) {
	return int(MaxFloat64(s))
}

// Argmin Return a single argmin value (first match)
func (s Int) Argmin() (argmin int) {
	return Argmin(s)
}

// Argmax Return a single argmax value (first match)
func (s Int) Argmax() (argmax int) {
	return Argmax(s)
}
func (s Int) ArgminAll() (argmin []int) {
	return ArgminAll(s)
}
func (s Int) ArgmaxAll() (argmax []int) {
	return ArgmaxAll(s)
}

func (s Int) Median() (mu float64) {
	return Median(s)
}
func (s Int) Quantile(p float64) (mu float64) {
	return Quantile(s, p)
}

// Var Unbiased variance calculation
// NOTE : Beware ! this is NOT sqrt(cov(x,x)), as it is divided by n-1 not by n
func (s Int) Var() (sigma2 float64) {
	return Var(s)
}

// Max
// Min
// Prod
// which.min(x)
// which.max
