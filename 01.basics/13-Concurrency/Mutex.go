package main

import (
	"fmt"
	"sync"
)

var (
	sum   int = 0
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add(10)
	}
	wg.Wait()
	fmt.Println("sum is ", sum)

}

func add(i int) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
