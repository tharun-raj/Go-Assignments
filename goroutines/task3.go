package main

import (
	"fmt"
	"sync"
)

func sayHello(message chan string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	message <- "Hello"
	m.Unlock()

}

func sayWorld(message chan string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	message <- "World"
	m.Unlock()

}

func main() {
	message := make(chan string)
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	wg.Add(2)
	go sayHello(message, wg, m)
	go sayWorld(message, wg, m)
	fmt.Println(<-message)
	fmt.Println(<-message)
	wg.Wait()
}
