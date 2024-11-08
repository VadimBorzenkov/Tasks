package main

import "fmt"

func ExInFirstButNorexInSecond(s1, s2 []string) []string {
	var res []string
	m := make(map[string]int)

	for _, value := range s2 {
		m[value]++
	}
	for _, value := range s1 {
		if count, found := m[value]; !found || count == 0 {
			res = append(res, value)
		} else {
			m[value]--
		}
	}
	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1", "apple", "apple"}
	slice2 := []string{"banana", "date", "fig", "apple", "hello"}
	res := ExInFirstButNorexInSecond(slice1, slice2)
	fmt.Println(res)
}
