package arbitrary

import "math/rand"

func Boolean(rng *rand.Rand) bool {
	return rand.Intn(2) == 0
}
