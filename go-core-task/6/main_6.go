package main

import(
	"math/rand"
	"fmt"
	"time"
	"sync"
)

func writer(ch chan int){
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < 3; i++{
		ch <- rand.Int()
	}
	close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for a := range ch{
		fmt.Println(a)
	}
}

func main(){
	ch  := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go writer(ch)
	go reader(ch, &wg)

	wg.Wait()
}