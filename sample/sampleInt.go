package sample

import (
	"math"
	"sort"
)

type Int []int

func NewInt(x []int) (s *Int) {
	s = new(Int)
	*s = Int(x)
	return
}

// implement sort
func (s Int) Len() int {
	return len(s)
}

func (s Int) Less(i, j int) bool {
	return s[i] <= s[j]
}
func (s Int) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
	return
}

func (s Int) Sum() (sum int) {
	sum = 0
	for _, v := range s {
		sum += v
	}
	return
}
func (s Int) Prod() (prod int) {
	prod = 1
	for _, v := range s {
		prod *= v
	}
	return
}

func (s Int) Length() (length int) {
	length = len(s)
	return
}

func (s Int) Mean() (mean float64) {
	length := s.Length()
	if length == 0 {
		return 0
	}
	mean = float64(s.Sum()) / float64(length)
	return
}

func (s Int) Min() (min int) {
	first := true
	for _, v := range s {
		if v < min {
			min = v
		}
		if first {
			first = false
			min = v
		}
	}
	return
}
func (s Int) Max() (max int) {
	first := true
	for _, v := range s {
		if v > max {
			max = v
		}
		if first {
			first = false
			max = v
		}
	}
	return
}

// Argmin Return a single argmin value (first match)
func (s Int) Argmin() (argmin int) {
	argmin = -1
	i := 0
	min := int(math.MaxInt64)
	for _, v := range s {
		{
		}
		if i == 0 || v < min {
			min = v
			argmin = i
		}
		i++
	}
	return
}

// Argmax Return a single argmax value (first match)
func (s Int) Argmax() (argmax int) {
	argmax = -1
	i := 0
	max := int(math.MinInt64)
	for _, v := range s {
		if i == 0 || v > max {
			max = v
			argmax = i
		}
		i++
	}
	return
}
func (s Int) ArgminAll() (argmin []int) {
	i := 0
	min := int(math.MaxInt64)
	for _, v := range s {
		if i == 0 || v <= min {
			min = v
			argmin = append(argmin, i)
		}
		i++
	}
	return
}
func (s Int) ArgmaxAll() (argmax []int) {
	i := 0
	max := int(math.MinInt64)
	for _, v := range s {
		if i == 0 || v >= max {
			max = v
			argmax = append(argmax, i)
		}
		i++
	}
	return
}

func (s Int) Median() (mu float64) {
	if s.Length() == 0 {
		panic("can't compute median on an empty sample")
	}
	sorted := s
	sort.Sort(sorted)
	length := len(sorted)
	if length%2 == 1 {
		mu = float64(sorted[length/2])
		return
	}
	mu = float64((sorted[length/2-1] + sorted[length/2])) / 2.
	return
}

func squareDiff(x Int) float64 {
	n := x.Length()
	mean := x.Mean()
	sum := 0.
	for i := 0; i < n; i++ {
		diff := float64(x[i]) - mean
		sum += diff * diff
	}
	return sum
}

// Var Unbiased variance calculation
// NOTE : Beware ! this is NOT sqrt(cov(x,x)), as it is divided by n-1 not by n
func (s Int) Var() (sigma2 float64) {
	n := s.Length()
	return squareDiff(s) / float64(n-1)
}

// Max
// Min
// Prod
// which.min(x)
// which.max
