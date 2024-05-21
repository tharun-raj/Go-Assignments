package main

import "fmt"

func some(msg chan string){
	msg <- "from function 1"

}

func another(msg chan string){
	msg <- "from function 2"
}

func main(){
	chan1 := make(chan string)
	chan2 := make(chan string)

	go some(chan1)
	go another(chan2)

	for i := 0; i < 2; i++{
		select{
			case val := <- chan1:
				fmt.Println(val)
			case val1 := <-chan2:
				fmt.Println(val1)
		}
	}
}