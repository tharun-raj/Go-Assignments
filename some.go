// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"


func f(n  chan int){
    n <- 10
    fmt.Println(<-n)
}

func main() {
    n := make(chan int)
     go f(n)
    fmt.Println(<-n, "main")
}