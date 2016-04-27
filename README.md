# rgo [![Build Status](https://travis-ci.org/eric-burel/rgo.svg?branch=master)](https://travis-ci.org/eric-burel/rgo)
Advanced random number simulation for Golang, inspired by R

Implements r*func*,d*func*,p*func* and q*func* functions for most common probabilistic laws.

The idea is to have a short, easy syntax for random variables generation and density calculation. 

This package is meant to be used for random scenarios generation (testing, benchmarks, user behaviour simulation, games and so on). 
However, if you want to use it for research purpose and need more precision/reliability/performance, do not hesitate to open issues (or better, pull requests) to propose improvements.

## Install
As simple as `go get` :
```
go get github.com/eric-burel/rgo
```
Then import the package in your code :
```go
import (
  "github.com/eric-burel/rgo"
)
```

## Complete documentation

[See godoc](https://godoc.org/github.com/eric-burel/rgo) for a list of all functionnalities.

An example with rand:

```go
package main

import (
  "github.com/eric-burel/rgo/rand"
)
x := rand.NewNorm(25.,1.)

todayTemperature := x.R()       // generate a random value, following a N(25,1) distributon
thisWeekTemperature := x.Rn(7)  // generate an array of 7 random values

f24 := x.D(24.)  // get f(24), density of this Gaussian distribution for x = 24

// we can then calculate likelihood of thisWeekTemperature
likelihood := 1
for _, t := range(thisWeekTemperature){
  likelihood *= x.D(t)
}

// Another example using standard definition of laws when they exist
y := rand.NewStdUnif()   // use classical definition, ie interval [0.,1.[
whiteNoise := y.R() // white noise is between 0. and 1.


```
An example of sample analysis :
```go
package main

import (
  "github.com/eric-burel/rgo/sample"
)

thisWeekTemperatures := []int{15,18,13,12,12,20,22}
x := sample.NewInt(thisWeekTemperatures) // generate a new rgo sample

avgTemperature := x.Mean()
variance := x.Var() // UNBIASED VARIANCE, with n-1 as the denominator (not n)
coldestDay := x.Argmin() // will return 3, first match of 12 degrees
coldestDays := x.ArgminAll() // will return []int{3,4}, days where temperatures was the coldest
hottestTemperature := x.Max() // 22

// same apply with float64
sprintDuration := []float64{10.009,10.001,9.887,9.129}
y := sample.NewFloat(sprintDuration)


...

```

## Random number generation

When possible, we use Golang standard lib math/rand to generate values.

So far, precision and quality of the randomness are NOT guaranteed. This package is used to generate random scenarios, for testing purposes or games, not for critical operations. 

Please avoid using this package to build a banking system or a flight autopilot (anyway, please don't code randomized flight autopilots).

## Density function
Density function calculation are implemented naively using their definitions. Optimal performances are not guaranteed.
They are tested against values given by R equivalent commands.

## Cumulative distribution function
Working on it...
## Quantile
Working on it...

## Edge cases
### Geometric law
Geometric law has two common definitions ([see wikipedia](https://en.wikipedia.org/wiki/Geometric_distribution)), depending on whether the support function includes 0 or not.

We chose the same definition as R, meaning __P(X=0) = p, with X following a geometric distribution G(p)__.
This behaviour may differ from what you are used to. For example in France, we prefer to start with k=1, meaning P(X=1) = p, so that k is the number of draws before a success.

### Empirical variance of a sample
The calculated variance ``x.Var()`` is unbiased, meaning the denominator is __n-1__, where __n__ is the size of the sample. This way, ``x.Var()`` is a good estimator of the actual variance of the probabilistic law under the sample.
It is different from the classical definition of the variance, with __n__ as the denominator, which is nice but biased.

Once again, this behaviour is the same as R.

# Roadmap
- [x] add most basic functions : bernouilli, uniform, binomial, poisson, normal, geom, exponential
- [ ] add other less common laws (Reverse Binomial *et al.*)
- [ ] implement quantile calculation
- [ ] implement cumulative density functions
- [ ] review naive implementations (mathematician needed)
- [ ] add error cases
- [ ] add multidimensionnal laws
- [ ] add more and more tests
- [ ] add performance tests
- [ ] add precision tests

# Contribution
All contributions are welcome. If you want to know more about Gontran project, please contact @eric-burel.
