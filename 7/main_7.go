package main

import (
	"fmt"
	"sync"
)


func mergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	
	output := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			merged <- val
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}


	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 7; i <= 9; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merged := mergeChannels(ch1, ch2, ch3)

	for val := range merged {
		fmt.Println(val)
	}
}
