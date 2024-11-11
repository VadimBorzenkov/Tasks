package main

import (
	"fmt"
	"sync"
)

func mergeChan(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergedCh <- id
				}
			}(ch, wg)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{11, 22, 33} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{111, 222, 333} {
			c <- num
		}
		close(c)
	}()

	for num := range mergeChan(a, b, c) {
		fmt.Println(num)
	}
}
