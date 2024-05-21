package main

import "fmt"

type WaitGroup struct{
	count int
}

func (wg *WaitGroup)Add(num int){
	wg.count += num
}

func (wg *WaitGroup)Done(){
	wg.count--
}

func (wg *WaitGroup)Wait(){
	for {
		if wg.count == 0{
			break
		}
	}
}

func looper(name string, wg *WaitGroup) {
	defer wg.Done()
	fmt.Println(name)
}

func main() {
	wg := WaitGroup{}
	wg.Add(3)
	go looper("route 1", &wg)
	go looper("route 2", &wg)
	go looper("route 3", &wg)
	wg.Wait()
}
