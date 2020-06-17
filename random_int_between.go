package helpers

import (
	"math/rand"
	"time"
)

func RandomIntBetween(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(max-min+1) + min
	return random
}
