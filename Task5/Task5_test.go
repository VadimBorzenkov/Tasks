package main

import (
	"testing"
)

// Проверка содержимого слайсов без учёта порядка
func sameElements(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	counts := make(map[int]int)
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

func TestIntersection(t *testing.T) {
	tests := []struct {
		slice1       []int
		slice2       []int
		expectedBool bool
		expected     []int
	}{
		{
			slice1:       []int{1, 2, 3, 4, 5},
			slice2:       []int{3, 4, 5, 6, 7},
			expectedBool: true,
			expected:     []int{3, 4, 5},
		},
		{
			slice1:       []int{10, 20, 30, 40},
			slice2:       []int{50, 60, 70},
			expectedBool: false,
			expected:     []int{},
		},
		{
			slice1:       []int{1, 2, 2, 3},
			slice2:       []int{2, 3, 3, 4},
			expectedBool: true,
			expected:     []int{2, 3},
		},
		{
			slice1:       []int{},
			slice2:       []int{1, 2, 3},
			expectedBool: false,
			expected:     []int{},
		},
		{
			slice1:       []int{5, 5, 5, 5},
			slice2:       []int{5, 5},
			expectedBool: true,
			expected:     []int{5, 5},
		},
	}

	for _, tt := range tests {
		resultBool, result := Intersection(tt.slice1, tt.slice2)

		// Проверка булевого значения
		if resultBool != tt.expectedBool {
			t.Errorf("Для входных данных %v и %v ожидалось %v, но получено %v", tt.slice1, tt.slice2, tt.expectedBool, resultBool)
		}

		// Проверка содержимого слайса без учёта порядка
		if !sameElements(result, tt.expected) {
			t.Errorf("Для входных данных %v и %v ожидалось %v, но получено %v", tt.slice1, tt.slice2, tt.expected, result)
		}
	}
}
