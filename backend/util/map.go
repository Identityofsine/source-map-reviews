package util

// AddMaps adds the values of b to a and returns a
func MergeMap(a map[string]any, b map[string]any) map[string]any {
	for key, value := range b {
		a[key] = value
	}
	return a
}
