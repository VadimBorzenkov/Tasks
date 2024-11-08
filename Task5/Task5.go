package main

import "fmt"

func Intersection(s1, s2 []int) (bool, []int) {
	var res []int
	m := make(map[int]int)
	for _, value := range s1 {
		m[value]++
	}
	for _, value := range s2 {
		if count, found := m[value]; count > 0 && found {
			res = append(res, value)
			m[value]--
		}
	}
	if len(res) == 0 {
		return false, res
	}
	return true, res
}

func main() {
	a := []int{1, 2, 2, 3}
	b := []int{2, 3, 3, 4}
	c, f := Intersection(a, b)
	fmt.Println(c, " ", f)
}
