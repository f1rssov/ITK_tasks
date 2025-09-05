package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {

	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	ch3 := make(chan int, 3)


	for i := 1; i <= 3; i++ {
		ch1 <- i
		ch2 <- i + 3
		ch3 <- i + 6
	}
	close(ch1)
	close(ch2)
	close(ch3)


	merged := mergeChannels(ch1, ch2, ch3)


	var result []int
	for val := range merged {
		result = append(result, val)
	}


	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, e := range expected {
		found := false
		for _, r := range result {
			if r == e {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Ожидался элемент %d, но он не найден в результате %v", e, result)
		}
	}

	if len(result) != len(expected) {
		t.Errorf("Ожидалась длина %d, но получена %d", len(expected), len(result))
	}
}
