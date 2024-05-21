package main

import (
	"fmt"
	"sync"
)

func looper(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go looper("route 1", wg)
	go looper("route 2", wg)
	go looper("route 3", wg)
	wg.Wait()
}
