package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func dog(dogChan chan bool, catChan chan bool) {
	i := 0
	for {
		select {
		case <-dogChan:
			fmt.Println("dog", i)
			i++
			catChan <- true
			break
		default:
			break
		}
	}
}

func cat(catChan chan bool, fishChan chan bool) {
	for {
		select {
		case <-catChan:
			fmt.Println("cat")
			fishChan <- true
			break
		default:
			break
		}
	}
}

func fish(fishChan chan bool, dogChan chan bool) {
	i := 0
	for {
		select {
		case <-fishChan:
			fmt.Println("fish")
			i++ // 计数，打印完之后就溜溜结束了。
			if i > 9 {
				wg.Done()
				return
			}
			dogChan <- true
			break
		default:
			break
		}
	}
}

func main() {
	dogChan, catChan, fishChan := make(chan bool), make(chan bool), make(chan bool)
	wg.Add(1)
	go dog(dogChan, catChan)
	go cat(catChan, fishChan)
	go fish(fishChan, dogChan)
	dogChan <- true // 记得这里进行启动条件，不然就没法启动了。
	wg.Wait()
}
