package main

import (
	"reflect"
	"testing"
)

// Проверка на корректность значений
func TestNaturalsToCubes(t *testing.T) {
	expected := []float64{0, 1, 8, 27, 64, 125, 216, 343, 512, 729, 1000}
	result := naturalsToCubes(10)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, но получено %v", expected, result)
	}
}
