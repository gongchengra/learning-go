package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum     int = 0
	rwmutex sync.RWMutex
	wg      sync.WaitGroup
)

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add(10)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		go readSum()
	}
	time.Sleep(2 * time.Second)
}

func readSum() {
	rwmutex.RLock()
	ret := sum
	rwmutex.RUnlock()
	fmt.Println("ret is ", ret)
}

func add(i int) {
	defer wg.Done()
	rwmutex.Lock()
	sum += i
	rwmutex.Unlock()
}
