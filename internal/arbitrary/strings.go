package arbitrary

import "math/rand"

func String(rng *rand.Rand, alphabet []rune, length int) string {
	characters := make([]rune, length)
	Fill[rune](rng, alphabet, characters)
	return string(characters)
}
