/**
 * Implements function to analyze a sample
 *
 * Any kind of sample can be implemented using Interface, provided it is able
 * to return its values as Float64, which is the highest precision we can achieve
 * with builtin types.
 * Then it can simply convert float64 in any types it like
 */
package sample

import "sort"

// Interface Minimum functions to be a sample
type Interface interface {
	sort.Interface
	Length() int
	GetValueFloat64(int) float64 // gets ith value of the sample, as a float64
}

// Sampler Interface for functions common to sampleFloat and sampleInt
type Sampler interface {
	Mean() float64
	Argmin() int
	Argmax() int
	ArgminAll() []int
	ArgmaxAll() []int
	Var() float64
	Median() float64
}

func Median(s Interface) (mu float64) {
	if s.Length() == 0 {
		panic("can't compute median on an empty sample")
	}
	sorted := s
	sort.Sort(sorted)
	length := sorted.Length()
	if length%2 == 1 {
		mu = sorted.GetValueFloat64(length / 2)
		return
	}
	mu = (sorted.GetValueFloat64(length/2-1) + sorted.GetValueFloat64(length/2)) / 2.
	return
}

// Calculate the sum of elements in the vector, as a float64
func sumFloat64(s Interface) (sum float64) {
	sum = 0
	for i := 0; i < s.Length(); i++ {
		sum += s.GetValueFloat64(i)
	}
	return
}

// Mean The mean of a sample
func Mean(s Interface) (mean float64) {
	length := s.Length()
	if length == 0 {
		return 0
	}
	mean = sumFloat64(s) / float64(length)
	return
}

func squareDiff(s Interface) float64 {
	mean := Mean(s)
	n := s.Length()
	sum := 0.
	for i := 0; i < n; i++ {
		diff := s.GetValueFloat64(i) - mean
		sum += diff * diff
	}
	return sum
}

// Var Variance of a sample
func Var(s Interface) (sigma2 float64) {
	n := s.Length()
	return squareDiff(s) / float64(n-1)
}

// Extremum Returns an extremum, given a comparison function
func extremum(s Interface, comp func(float64, float64) bool) float64 {
	l := s.Length()
	switch l {
	case 0:
		panic("can not calculate min of an empty sample")
	case 1:
		return s.GetValueFloat64(0)
	default:
		extremum := s.GetValueFloat64(0)
		for i := 1; i < l; i++ {
			v := s.GetValueFloat64(i)
			if comp(v, extremum) {
				extremum = v
			}
		}
		return extremum
	}

}

// MinFloat64 Minimal element of a sample, as a Float64
func MinFloat64(s Interface) (min float64) {
	return extremum(s, func(x, y float64) bool { return x < y })
}

// MaxFloat64 Minimal element of a sample, as a Float64
func MaxFloat64(s Interface) (min float64) {
	return extremum(s, func(x, y float64) bool { return x > y })
}

// SumFloat64 Sum of the elements, as a float64
func SumFloat64(s Interface) (sum float64) {
	sum = 0.
	for i := 0; i < s.Length(); i++ {
		sum += s.GetValueFloat64(i)
	}
	return
}

// ProdFloat64 Return products of the interface element, as a float64
func ProdFloat64(s Interface) (prod float64) {
	prod = 1.
	for i := 0; i < s.Length(); i++ {
		v := s.GetValueFloat64(i)
		prod *= v
	}
	return
}

// argextremum Calculates argmin or argmax, depending on the comp function
func argextremum(s Interface, comp func(float64, float64) bool) (arg int) {
	l := s.Length()
	switch l {
	case 0:
		panic("can not calculate arg-min or arg-max of an empty sample")
	case 1:
		return 0
	default:
		extremum := extremum(s, comp)
		arg = 0
		for i := 1; i < l; i++ {
			v := s.GetValueFloat64(i)
			if v == extremum {
				arg = i
				break
			}
		}
		return arg
	}
}

// argextremum Calculates argmin or argmax (not strict), depending on the comp function
func argextremumAll(s Interface, comp func(float64, float64) bool) (arg []int) {
	l := s.Length()
	switch l {
	case 0:
		panic("can not calculate arg-min or arg-max of an empty sample")
	case 1:
		return []int{0}
	default:
		extremum := extremum(s, comp)
		arg = []int{}
		for i := 0; i < l; i++ {
			v := s.GetValueFloat64(i)
			if v == extremum {
				arg = append(arg, i)
			}
		}
		return arg
	}
}

// Argmin Index of a minimal element of s (first found)
func Argmin(s Interface) int {
	return argextremum(s, func(x, y float64) bool { return x < y })
}

// Argmax Index of a maximal element of s (first found)
func Argmax(s Interface) int {
	return argextremum(s, func(x, y float64) bool { return x > y })
}

// ArgminAll Indexes of the minimal elements of s
func ArgminAll(s Interface) []int {
	return argextremumAll(s, func(x, y float64) bool { return x < y })
}

// ArgmaxAll Indexes of the maximal elements of s
func ArgmaxAll(s Interface) []int {
	return argextremumAll(s, func(x, y float64) bool { return x > y })
}
