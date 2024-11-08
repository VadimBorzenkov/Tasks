package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	res := make(map[string]int)
	for key, value := range m.data {
		res[key] = value
	}
	return res
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.data[key]
	return value, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("1", 1)
	m.Add("2", 2)
	m.Add("3", 3)
	m.Add("4", 4)
	fmt.Println("1)")
	for _, value := range m.data {
		fmt.Println(value)
	}
	fmt.Println("2)")

	m.Remove("1")
	for _, value := range m.data {
		fmt.Println(value)
	}

	fmt.Println("3)")
	cp := m.Copy()
	for _, value := range cp {
		fmt.Println(value)
	}

	fmt.Println("4)")
	fmt.Println("Cуществует: ", m.Exists("3"))
	fmt.Println("Не существует: ", m.Exists("1"))

	fmt.Println("5)")
	value1, f1 := m.Get("3")
	value2, f2 := m.Get("1")
	fmt.Println("Cуществующее значение: ", value1, f1)
	fmt.Println("Не существующее значение: ", value2, f2)

}
