package internal

// TODO: Unit test
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, el := range slice {
		if predicate(el) {
			result = append(result, el)
		}
	}
	return result
}
