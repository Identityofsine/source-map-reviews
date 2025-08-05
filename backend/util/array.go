package util

func ToGenericArray[T any](items ...T) []interface{} {
	if len(items) == 0 {
		return nil
	}
	result := make([]interface{}, len(items))
	for i, item := range items {
		result[i] = item
	}
	return result
}

func FlatList[T any](items [][]T) []T {
	if len(items) == 0 {
		return nil
	}
	var result []T
	for _, item := range items {
		result = append(result, item...)
	}
	return result
}

func PtrArray[T any](items ...T) []*T {
	if len(items) == 0 {
		return nil
	}
	result := make([]*T, len(items))
	for i, item := range items {
		result[i] = &item
	}
	return result
}
