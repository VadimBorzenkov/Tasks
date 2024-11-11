package main

import "testing"

func TestGenRand(t *testing.T) {
	// Проверка правильного количества чисел
	n := 10
	ch := genRand(n)

	// Проверяем, что канал содержит ровно n чисел
	count := 0
	for val := range ch {
		count++
		if val < 0 || val >= n {
			t.Errorf("Generated value %d is out of range [0, %d)", val, n)
		}
	}
	if count != n {
		t.Errorf("Expected %d values, but got %d", n, count)
	}

	// Проверка, что канал закрывается после завершения генерации
	ch = genRand(n)
	for range ch {
	}
	_, open := <-ch
	if open {
		t.Errorf("Expected channel to be closed after generating %d values", n)
	}
}
