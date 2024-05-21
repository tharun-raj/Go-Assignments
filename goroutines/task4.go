package main

import (
	"fmt"
)

func printTen(message chan int, done chan bool) {
	for i := 1; i <= 10; i++ {
		message <- i
	}
	done <- true
	close(message)
}

func main() {
	message := make(chan int)
	done := make(chan bool)
	go printTen(message, done)
Loop:
	for {
		select {
		case val := <-message:
			fmt.Println(val)
		case <-done:
			break Loop
		}
	}
}
