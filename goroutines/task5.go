package main

import "fmt"

func printChannel(channel chan int){
	for i := 0; i < cap(channel); i++{
		fmt.Println(<-channel)
	}
	close(channel)
}

func main(){
	channel := make(chan int, 3)
	for i := 1; i <= 3; i++{
		channel <- i
	}
	printChannel(channel)
}
