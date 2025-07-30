package util

func ToGenericArray[T any](items ...T) []interface{} {
	if len(items) == 0 {
		return nil
	}
	result := make([]interface{}, 0, len(items))
	for i, item := range items {
		result[i] = item
	}
	return result
}
