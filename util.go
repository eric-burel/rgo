package rgo

// IsProba Checks whether a float64 represent a probability or not
func IsProba(p float64) (bool){
    return (p >= 0. && p <= 1.0)
}

type intArr []int
type float64Arr []float64

// Sum Returns the sum of an int array
// TODO Refactor me into one sum function using interface or slices
func (array intArr) Sum() (s int){
    s = 0
    for _, v := range(array){
        s += v
    }
    return
}

// Sum Returns the sum of a float64 array
// TODO Refactor me into one sum function using interface or slices
func (array float64Arr) Sum() (s float64){
    s = 0.
    for _, v := range(array){
        s += v
    }
    return
}
