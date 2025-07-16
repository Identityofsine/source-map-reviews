package util

func GroupBy[T any](items []T, key func(T) string) map[string][]T {
	grouped := make(map[string][]T)
	for _, item := range items {
		grouped[key(item)] = append(grouped[key(item)], item)
	}
	return grouped
}

func GroupIntoLists[T any](items []T, key func(T) string) [][]T {
	grouped := make(map[string][]T)
	for _, item := range items {
		grouped[key(item)] = append(grouped[key(item)], item)
	}
	return values(grouped)
}

func values[T any](grouped map[string][]T) [][]T {
	values := make([][]T, 0, len(grouped))
	for _, group := range grouped {
		values = append(values, group)
	}
	return values
}
