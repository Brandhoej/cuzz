package arbitrary

import "math/rand"

func From[T any](rng *rand.Rand, collection []T) T {
	return collection[rand.Intn(len(collection))]
}

func Fill[T any](rng *rand.Rand, alphabet, data []T) {
	for i := range data {
		data[i] = From[T](rng, alphabet)
	}
}

func Shuffel[T any](rng *rand.Rand, data []T) {
	rng.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}
