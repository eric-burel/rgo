package rand

import(
    "errors"
)


// Lawer Implements a probalistic law
type Lawer interface{
    // Generate
    Rn(n int) []interface{}
    R() interface{}
    // probability density function
    D(value interface{}) float64
    // cumulative probability density function
    P(value interface{}) float64
    // quantile
    Q(value float64) float64
}

// Discreter Implements a discrete law
type Discreter interface{
    Rn(length int) []int
    R() int
    D(value int) float64
    P(value int) float64
}

// Continuouser Implements a continous law
type Continuouser interface{
    Rn(length int) []float64
    R() float64
    D(value float64) float64
    P(value float64) float64
}

// Continuous A Continuous law
type Continuous struct{
    Continuouser
}
// Discrete A discret law
type Discrete struct{
    Discreter
}

// Rn Generate n draws of a continuous random variable
func (c Continuous) Rn(n int) (res []float64){
    if (n < 0){
        panic(errors.New("n must be positive"))
    }
    res = make([]float64,n)
    for i := 0; i < n; i++{
        res[i] = c.R()
    }
    return
}
// Rn Generate n draws of a discrete random variable
func (d Discrete) Rn(n int) (res []int) {
    if (n < 0){
        panic(errors.New("n must be positive"))
    }
    res = make([]int,n)
    for i := 0; i < n; i++{
        res[i] = d.R()
    }
    return
}
