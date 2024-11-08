package main

import (
	"testing"
)

// Тест на сравнение слайсов без учёта порядка
func sameElements(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	counts := make(map[string]int)
	for _, v := range slice1 {
		counts[v]++
	}
	for _, v := range slice2 {
		if counts[v] == 0 {
			return false
		}
		counts[v]--
	}
	return true
}

func TestExInFirstButNorexInSecond(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			slice1:   []string{"apple", "apple", "banana"},
			slice2:   []string{"banana"},
			expected: []string{"apple", "apple"},
		},
		{
			slice1:   []string{"a", "b", "c", "a"},
			slice2:   []string{"b", "a"},
			expected: []string{"c", "a"},
		},
		{
			slice1:   []string{},
			slice2:   []string{"x", "y"},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		result := ExInFirstButNorexInSecond(tt.slice1, tt.slice2)
		if !sameElements(result, tt.expected) {
			t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
		}
	}
}
