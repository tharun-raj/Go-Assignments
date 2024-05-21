package main

import "fmt"

func Greet(message chan string) {
	message <- "hello world"
}

func main() {
	message := make(chan string)
	go Greet(message)
	fmt.Println(<-message)
}
