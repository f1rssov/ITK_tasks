package main

import (
	"fmt"
	"time"
)


type CustomWaitGroup struct {
	done chan struct{}
	cnt  int
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		done: make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(n int) {
	wg.cnt += n
}

func (wg *CustomWaitGroup) Done() {
	wg.done <- struct{}{}
}

func (wg *CustomWaitGroup) Wait() {
	for i := 0; i < wg.cnt; i++ {
		<-wg.done
	}
}

func worker(id int, wg *CustomWaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(100*id))
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	wg := NewCustomWaitGroup()
	numWorkers := 3
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
