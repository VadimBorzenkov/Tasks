package main

import (
	"reflect"
	"testing"
)

func TestNewSlice(t *testing.T) {
	res := newSlice()

	if len(res) != 10 {
		t.Errorf("ожидаемая длина слайса 10, но получили %d", len(res))

		for _, value := range res {
			if value < 0 || value > 99 {
				t.Errorf("найдено значение вне диапазона [0, 99]: %d", value)
			}
		}
	}
}

func TestSliceExample(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{2, 4, 6, 8}
	res := sliceExample(sl)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Ожидаемый результат %v, но получили %v", expected, res)
	}
}

func TestAddElements(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := addElements(sl, 10)
	if len(res) != len(sl)+1 {
		t.Errorf("Длина нового массива должна быть на 1 больше длины исходного")
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Ожидаемый результат %v, но получили %v", expected, res)
	}
}

func TestCopySlice(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := copySlice(sl)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Ожидаемый результат %v, но получили %v", expected, res)
	}
}

func TestRemoveElement(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{2, 3, 4, 5, 6, 7, 8}
	res := removeElement(sl, 0)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Ожидаемый результат %v, но получили %v", expected, res)
	}
}
