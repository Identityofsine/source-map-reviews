package util

// AddMaps adds the values of b to a and returns a
func MergeMap(a map[string]any, b map[string]any) map[string]any {
	for key, value := range b {
		a[key] = value
	}
	return a
}

// MapBy will mutate the input array and then group them by the key; this is useful for grouping items by a specific key
func MapBy[I any, T any](a []I, k func(I) string, i func(I) T) map[string]T {
	newMap := make(map[string]T)
	for _, item := range a {
		key := k(item)
		newMap[key] = i(item) // mutate the input map
	}
	return newMap
}

// Map will mutate the input map and return the input map
func Map[I any, T any](a []I, f func(I) T) []T {
	newArr := make([]T, len(a))
	for i, item := range a {
		newArr[i] = f(item) // mutate the input array
	}
	return newArr
}

func MapToMap[I any, T any](a map[string]I, f func(I) T) map[string]T {
	newMap := make(map[string]T)
	for key, value := range a {
		newMap[key] = f(value) // mutate the input map
	}
	return newMap
}

// MapValues will mutate the input map and return a new map with the values transformed by the function f
// f is a function that takes an input of type I and returns a value of type T
// This is a "valueGetter"
func GetMapValues[I any, T any](a map[string]I, f func(I) T) []T {
	newArr := make([]T, 0)
	for _, item := range a {
		newArr = append(newArr, f(item)) // mutate the input array
	}

	return newArr
}

func GetMapKeys[I comparable, T any](a map[I]T) []I {

	newArr := make([]I, 0)
	for key, _ := range a {
		newArr = append(newArr, key)
	}

	return newArr
}
