package main

import(
	"testing"
	"reflect"
)

func TestSliceExample(t *testing.T){
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "четные числа",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "только нечётные",
			input:    []int{1, 3, 5, 7},
			expected: nil,
		},
		{
			name:     "только чётные",
			input:    []int{2, 4, 6, 8},
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceExample(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}	
}

func TestAddElements(t *testing.T){
		tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "1 Добавление",
			input:    []int{1},
			expected: []int{1, 6},
		},
		{
			name:     "2 Добавление",
			input:    []int{1, 3,},
			expected: []int{1, 3, 6},
		},
		{
			name:     "3 Добавление",
			input:    []int{2, 4, 6, 8},
			expected: []int{2, 4, 6, 8, 6},
		},
		{
			name:     "4 Добавление в пустой слайс",
			input:    []int{},
			expected: []int{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addElements(tt.input, 6)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}	
}

func  TestCopySlice(t *testing.T){
			tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Копирование",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Копирование",
			input:    []int{1, 3},
			expected: []int{1, 3},
		},
		{
			name:     "Копирование пустого",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copySlice(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
			tt.input = append(tt.input, 3)
			if reflect.DeepEqual(result, tt.input) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}	
}

func  TestRemoveElement(t *testing.T){
			tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Удаление 2-ого числа, из слайса длинны 1",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Удаление 2-ого числа",
			input:    []int{1, 3},
			expected: []int{1},
		},
		{
			name:     "Удаление из пустого",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElement(tt.input, 2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}	
}