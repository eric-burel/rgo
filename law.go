package rgo

import(
    "math"
)

// Probability A float belonging to [0,1]
type Probability float64

// Law Implements a probalistic law
type Law interface{
    // Generate
    Rn(n int) []interface{}
    R() interface{}
    // probability density function
    D(value interface{}) Probability
    // cumulative probability density function
    P(value interface{}) Probability
    // quantile
    Q(value Probability) float64
}

// Discrete Implements a discrete law
type Discrete interface{
    Rn(length int) []int
    R() int
    D(value int) Probability
    P(value int) Probability
}

// Continuous Implements a continous law
type Continuous interface{
    Rn(length float64) []int
    R() int
    D(value float64) Probability
    P(value float64) Probability
}
