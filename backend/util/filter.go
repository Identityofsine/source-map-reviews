package util

func Filter[I any](a []I, f func(I) bool) []I {
	newArr := make([]I, 0)
	for _, item := range a {
		if f(item) {
			newArr = append(newArr, item) // mutate the input array
		}
	}
	return newArr
}
