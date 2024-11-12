package main

import (
	"sync"
	"testing"
	"time"
)

func TestSem(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}

	s := NewSem(len(numbers))

	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()
			s.Add(1)
		}(num)
	}

	// Устанавливаем таймаут, чтобы избежать бесконечного ожидания
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Все горутины завершились успешно
	case <-time.After(2 * time.Second):
		t.Error("Timeout waiting for goroutines to finish")
	}

	// Проверка, что Done корректно освобождает семафор
	s.Done(len(numbers))
	select {
	case s <- struct{}{}:
		// Если Done корректно освобождает семафор, то семафор не должен блокировать
	default:
		t.Errorf("Expected semaphore to be empty after Done")
	}
}
