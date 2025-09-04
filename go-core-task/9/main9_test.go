package main

import (
	"reflect"
	"testing"
)

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8, 5)
	out := make(chan float64, 5)

	// пишем тестовые данные в канал
	for i := uint8(1); i <= 5; i++ {
		in <- i
	}
	close(in)

	// запускаем конвейер
	go cubePipeline(in, out)

	// собираем результаты
	var result []float64
	for val := range out {
		result = append(result, val)
	}

	// ожидаемые кубы 1..5
	expected := []float64{1, 8, 27, 64, 125}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
