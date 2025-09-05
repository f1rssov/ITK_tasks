package main

import (
	"fmt"
)

func cubePipeline(in <-chan uint8, out chan<- float64) {
	for num := range in {
		f := float64(num)
		out <- f * f * f
	}
	close(out)
}

func main(){
	in := make(chan uint8)
	out := make(chan float64)

	
	go cubePipeline(in, out)


	go func() {
		for i := uint8(1); i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	// читаем результаты из выходного канала
	for result := range out {
		fmt.Println(result)
	}
}
