package main

import (
	"fmt"
	"sync"
	"time"
)

func looper(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		time.Sleep(1)
		fmt.Println(name+" : ", i)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go looper("route 1", wg)
	go looper("route 2", wg)
	go looper("route 3", wg)
	wg.Wait()
}
