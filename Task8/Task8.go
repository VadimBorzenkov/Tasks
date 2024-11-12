package main

import "fmt"

type sem chan struct{}

func NewSem(n int) sem {
	return make(sem, n)
}

func (s sem) Add(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s sem) Done(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	s := NewSem(len(numbers))

	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			s.Add(1)
		}(num)
	}

	s.Done(len(numbers))
}
