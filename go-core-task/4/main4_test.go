package main

import (
	"reflect"
	"testing"
)

func TestMakeDifference(t *testing.T) {
	tests := []struct {
		name     string
		s1       []string
		s2       []string
		expected []string
	}{
		{
			name:     "обычные элементы",
			s1:       []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			s2:       []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "пустой второй слайс",
			s1:       []string{"a", "b"},
			s2:       []string{},
			expected: []string{"a", "b"},
		},
		{
			name:     "пустой первый слайс",
			s1:       []string{},
			s2:       []string{"x", "y"},
			expected: nil,
		},
		{
			name:     "нет общих элементов",
			s1:       []string{"a", "b"},
			s2:       []string{"x", "y"},
			expected: []string{"a", "b"},
		},
		{
			name:     "все элементы второго слайса присутствуют в первом",
			s1:       []string{"a", "b", "c"},
			s2:       []string{"a", "b", "c"},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := makeDifference(tt.s1, tt.s2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}