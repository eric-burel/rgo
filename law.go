package rgo



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

// Discrete Implements a discrete law
type Discrete interface{
    Rn(length int) []int
    R() int
    D(value int) float64
    P(value int) float64
}

// Continuous Implements a continous law
type Continuous interface{
    Rn(length float64) []int
    R() int
    D(value float64) float64
    P(value float64) float64
}
