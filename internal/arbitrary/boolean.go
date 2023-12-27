package arbitrary

import "math/rand"

func Boolean(rng *rand.Rand) bool {
	return rand.Int63()&1 == 0
}
