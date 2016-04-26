package rand

import (
	"math/rand"
	"time"
)

var seeded = false

// Generate a new seed
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	seeded = true
}
