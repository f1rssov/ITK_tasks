package main

import (
	"reflect"
	"testing"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	// ==== Add и Exists ====
	m.Add("one", 1)
	m.Add("two", 2)

	if !m.Exists("one") {
		t.Error("Ожидалось, что ключ 'one' существует")
	}
	if !m.Exists("two") {
		t.Error("Ожидалось, что ключ 'two' существует")
	}
	if m.Exists("three") {
		t.Error("Ожидалось, что ключ 'three' не существует")
	}

	// ==== Get ====
	val, ok := m.Get("one")
	if !ok || val != 1 {
		t.Errorf("Ожидалось: 1, true; Получено: %d, %v", val, ok)
	}

	val, ok = m.Get("three")
	if ok {
		t.Errorf("Ожидалось: false для ключа 'three'; Получено: %v", ok)
	}

	// ==== Remove ====
	m.Remove("one")
	if m.Exists("one") {
		t.Error("Ожидалось, что ключ 'one' был удалён")
	}

	// ==== Copy ====
	copyMap := m.Copy() 
	expected := map[string]int{"two": 2}
	if !reflect.DeepEqual(copyMap, expected) {
		t.Errorf("Ожидалось: %v; Получено: %v", expected, copyMap)
	}

	// Проверяем, что изменения в оригинале не влияют на копию
	m.Add("four", 4)
	if _, ok := copyMap["four"]; ok {
		t.Error("Копия не должна изменяться при изменении оригинала")
	}
}
