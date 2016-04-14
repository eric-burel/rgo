# rgo
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


## Random number generation

When possible, we use Golang standard lib math/rand to generate values.

So far, precision and quality of the randomness are NOT guaranteed. This package is used to generate random scenarios, for testing purposes or games, not for critical operations. 

Please avoid using this package to build a banking system or a flight autopilot (anyway, please don't code randomized flight autopilots).

## Density function
Density function calculation are implemented naively using their definitions. Optimal performances are not guaranteed.

## Cumulative distribution function
Working on it...
## Quantile
Working on it...

# Roadmap
- [ ] add most common probability laws (Gaussian, Poisson and so on)
- [ ] add other less common laws (Reverse Binomial *et al.*)
- [ ] implement quantile calculation
- [ ] implement cumulative density function
- [ ] review naive implementations (mathematician needed)
- [ ] add error cases
- [ ] add multidimensionnal laws
- [ ] add performance tests
- [ ] add precision tests

# Contribution
All contributions are welcome. If you want to know more about Gontran project, please contact @eric-burel.
