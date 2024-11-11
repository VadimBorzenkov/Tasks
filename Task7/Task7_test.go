package main

import (
	"testing"
	"time"
)

func TestMergeChan(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	go func() {
		defer close(ch2)
		ch2 <- 4
		ch2 <- 5
	}()

	go func() {
		defer close(ch3)
		ch3 <- 6
		ch3 <- 7
		ch3 <- 8
	}()

	mergedCh := mergeChan(ch1, ch2, ch3)

	// Создаем карту, чтобы отслеживать все полученные значения
	expectedValues := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true}
	receivedValues := make(map[int]bool)

	for val := range mergedCh {
		receivedValues[val] = true
	}

	// Проверяем, что получены все ожидаемые значения
	for val := range expectedValues {
		if !receivedValues[val] {
			t.Errorf("Expected value %d not received from merged channel", val)
		}
	}

	// Проверяем, что лишних значений нет
	if len(receivedValues) != len(expectedValues) {
		t.Errorf("Received unexpected number of values: got %d, want %d", len(receivedValues), len(expectedValues))
	}

	// Проверка на закрытие канала
	select {
	case _, ok := <-mergedCh:
		if ok {
			t.Error("Expected merged channel to be closed")
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout: Expected merged channel to be closed")
	}
}
