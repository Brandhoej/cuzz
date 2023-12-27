package arrangement

func Cartesian[T any](values [][]T) [][]T {
	if len(values) == 0 {
		return [][]T{}
	}

	var helper func(current []T, depth int)
	result := [][]T{}

	helper = func(current []T, depth int) {
		if depth == len(values) {
			result = append(result, append([]T{}, current...))
			return
		}

		for _, value := range values[depth] {
			helper(append(current, value), depth+1)
		}
	}

	helper([]T{}, 0)
	return result
}
