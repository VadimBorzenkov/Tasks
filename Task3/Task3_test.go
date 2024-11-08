package main

import "testing"

func TestAddElement(t *testing.T) {
	m := NewStringIntMap()
	m.Add("test", 1)

	value, exists := m.Get("test")
	if !exists {
		t.Errorf("Ключ 'test' не найден")
	} else if value != 1 {
		t.Errorf("Оиждалось 1, а получено %d", value)
	}
}

func TestRemoveElement(t *testing.T) {
	m := StringIntMap{map[string]int{"1": 1}}
	m.Remove("1")
	if _, exists := m.Get("1"); exists {
		t.Errorf("Элемент не был удалён")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Add("key2", 200)

	copyMap := m.Copy()

	// Проверяем, что все элементы в копии совпадают с оригиналом
	for key, value := range m.data {
		if copyMap[key] != value {
			t.Errorf("Значение для ключа '%s' должно быть %d, но получено %d", key, value, copyMap[key])
		}
	}

	// Изменяем копию и проверяем, что оригинал не изменился
	copyMap["key1"] = 300
	if m.data["key1"] == 300 {
		t.Errorf("Изменение в копии не должно влиять на оригинальную карту")
	}

	// Проверяем, что добавление нового элемента в копию не влияет на оригинал
	copyMap["newKey"] = 400
	if _, exists := m.data["newKey"]; exists {
		t.Errorf("Добавление нового элемента в копию не должно добавлять его в оригинал")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Add("key2", 200)

	// Проверка существующего ключа
	if !m.Exists("key1") {
		t.Errorf("Метод работает неверно, потому что элемент 'key1' находится в структуре")
	}

	// Проверка отсутствующего ключа
	if m.Exists("key3") {
		t.Errorf("Метод работает неверно, потому что элемента 'key3' нет в структуре")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Add("key2", 200)

	// Тест существующего ключа
	value, ok := m.Get("key1")
	if !ok {
		t.Errorf("Ожидалось, что ключ 'key1' существует, но метод вернул false")
	}
	if value != 100 {
		t.Errorf("Ожидалось значение 100 для ключа 'key1', но получено %d", value)
	}

	// Тест несуществующего ключа
	value, ok = m.Get("key3")
	if ok {
		t.Errorf("Ожидалось, что ключ 'key3' отсутствует, но метод вернул true")
	}
	if value != 0 {
		t.Errorf("Ожидалось значение 0 для отсутствующего ключа 'key3', но получено %d", value)
	}
}
