package main

import (
	"testing"
	"time"
)


func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()
	numWorkers := 3
	wg.Add(numWorkers)

	doneSignals := make(chan int, numWorkers)


	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			// имитация работы
			time.Sleep(10 * time.Millisecond)
			doneSignals <- id
		}(i)
	}

	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)

	if elapsed < 10*time.Millisecond {
		t.Errorf("Wait завершился слишком рано, должно было подождать всех Done")
	}


	results := make(map[int]bool)
	for i := 0; i < numWorkers; i++ {
		id := <-doneSignals
		results[id] = true
	}

	for i := 0; i < numWorkers; i++ {
		if !results[i] {
			t.Errorf("Горутина %d не завершилась", i)
		}
	}
}


func TestAddDoneDirect(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(2)

	doneCount := 0
	go func() {
		wg.Done()
		doneCount++
	}()
	go func() {
		wg.Done()
		doneCount++
	}()

	wg.Wait()

	if doneCount != 2 {
		t.Errorf("Ожидалось doneCount=2, получено %d", doneCount)
	}
}
