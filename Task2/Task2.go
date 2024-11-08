package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newSlice() []int {
	rand.Seed(time.Now().UnixNano())
	var res []int
	for i := 0; i < 10; i++ {
		res = append(res, rand.Intn(100))
	}
	return res
}

func sliceExample(sl []int) []int {
	var res []int
	for _, value := range sl {
		if value%2 == 0 {
			res = append(res, value)
		}
	}
	return res
}

func addElements(sl []int, value int) []int {
	res := make([]int, len(sl), len(sl)*2)
	copy(res, sl)
	res = res[:len(sl)+1]
	res[len(sl)] = value
	return res
}

func copySlice(sl []int) []int {
	res := make([]int, len(sl))
	for i := range res {
		res[i] = sl[i]
	}
	return res
}

func removeElement(sl []int, index int) []int {
	if index < 0 || index > len(sl) {
		return sl
	}
	return append(sl[:index], sl[index+1:]...)
}

func main() {
	sl := newSlice()
	fmt.Println("Исходный ", sl)

	ex := sliceExample(sl)
	fmt.Println("Только четные:", ex)

	ap := addElements(sl, 10)
	fmt.Println("Добавлен элемент", ap)

	cp := copySlice(sl)
	sl[0] = 1
	fmt.Println("Скопированный ", cp, "\nИсходный", sl)

	rm := removeElement(sl, 3)
	fmt.Println("Исходный с удаленным элементом по индексу 3 ", rm)
}
