package main

import (
	"testing"
)

func TestMakeIntersection(t *testing.T) {
	tests := []struct {
		name     string
		s1       []int
		s2       []int
		expectedOk bool
		expectedRes []int
	}{
		{
			name: "обычные пересечения",
			s1: []int{65, 3, 58, 678, 64},
			s2: []int{64, 2, 3, 43},
			expectedOk: true,
			expectedRes: []int{3, 64},
		},
		{
			name: "нет пересечений",
			s1: []int{1, 2, 3},
			s2: []int{4, 5, 6},
			expectedOk: false,
			expectedRes: []int{},
		},
		{
			name: "пустой первый слайс",
			s1: []int{},
			s2: []int{1, 2},
			expectedOk: false,
			expectedRes: []int{},
		},
		{
			name: "пустой второй слайс",
			s1: []int{1, 2},
			s2: []int{},
			expectedOk: false,
			expectedRes: []int{},
		},
		{
			name: "оба пустых слайса",
			s1: []int{},
			s2: []int{},
			expectedOk: false,
			expectedRes: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, res := makeIntersection(tt.s1, tt.s2)
			
			if ok != tt.expectedOk {
				t.Errorf("Ожидалось ok = %v, получено %v", tt.expectedOk, ok)
			}
			
			if len(res) != len(tt.expectedRes) {
				t.Errorf("Ожидалось %v элементов, получено %v элементов", len(tt.expectedRes), len(res))
			} else {
				resMap := make(map[int]struct{})
				for _, v := range res {
					resMap[v] = struct{}{}
				}
				for _, v := range tt.expectedRes {
					if _, ok := resMap[v]; !ok {
						t.Errorf("Ожидалось, что элемент %v будет в пересечении", v)
					}
				}
			}
		})
	}
}
