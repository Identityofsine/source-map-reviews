package db

import (
	"fmt"
	"strings"
)

func Placeholders(n int) string {
	if n <= 0 {
		return ""
	}
	placeholders := make([]string, n)
	for i := 0; i < n; i++ {
		placeholders[i] = fmt.Sprintf("$%d", i+1) // PostgreSQL style
	}
	return strings.Join(placeholders, ", ")
}
