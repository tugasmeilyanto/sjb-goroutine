package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.waitGroup

// 1. add
// 2. wait
// 3. done

// channel

// var channel <- true
// variabel := true

func Task(wg *sync.WaitGroup, msg string, done chan bool) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println(msg)
	done <- true
}

func main() {
	wg := sync.WaitGroup{}
	done := make(chan bool) // true

	for i := 0; i < 20; i++ {
		teks := fmt.Sprintf("task %d", i+1)
		wg.Add(1)
		go Task(&wg, teks, done)
		<-done
	}

	wg.Wait()
	close(done)
}
