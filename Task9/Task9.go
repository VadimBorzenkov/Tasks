package main

import (
	"fmt"
	"math"
)

func naturalsToCubes(n int) []float64 {
	naturals := make(chan uint8)
	cubes := make(chan float64)
	var results []float64

	go func() {
		for i := 0; i <= n; i++ {
			naturals <- uint8(i)
		}
		close(naturals)
	}()

	go func() {
		for i := range naturals {
			cubes <- math.Pow(float64(i), 3)
		}
		close(cubes)
	}()

	for i := range cubes {
		results = append(results, i)
	}
	return results
}

func main() {
	results := naturalsToCubes(11)

	for _, value := range results {
		fmt.Println(value)
	}
}
