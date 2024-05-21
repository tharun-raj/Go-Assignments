package main

import (
	"fmt"
	"sync"
)

func sayHello(message chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	message <- "Hello"
}

func sayWorld(message chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	message <- "World"
}

func main() {
	message := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go sayHello(message, wg)
	go sayWorld(message, wg)
	fmt.Println(<-message)
	fmt.Println(<-message)
	wg.Wait()
}
