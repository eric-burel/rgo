package rgo

// IsProba Checks whether a float represent a probability or not
func IsProba(p float64) (bool){
    return (p >= 0. && p <= 1.0)
}
