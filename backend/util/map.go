package util

// AddMaps adds the values of b to a and returns a
func MergeMap(a map[string]any, b map[string]any) map[string]any {
	for key, value := range b {
		a[key] = value
	}
	return a
}

// Map will mutate the input map and return the input map
func Map[I any, T any](a []I, f func(I) T) []T {
	newArr := make([]T, len(a))
	for i, item := range a {
		newArr[i] = f(item) // mutate the input array
	}
	return newArr
}
