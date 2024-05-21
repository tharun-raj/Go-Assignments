package main

import "fmt"

type Counter struct{
	count int
}

func (counter Counter) add()  Counter{ 
    counter.count += 1
	return counter
}

func (counter Counter) subtract()  Counter{ 
    counter.count -= 1
	return counter
}

func (counter Counter) display() { 
    fmt.Println(counter.count)
}

func (counter Counter) reset()  Counter{ 
    counter.count = 0
	return counter
}

func calculateAverage(numbers []int) int {
	sum := 0
	//use _ instead of i because _  if you use i instead of _,
	//it creates a new variable i which holds the index of the current iteration.
	//If you're not using this index variable, Go will complain because you've declared a variable
	for _, v := range numbers {
		sum += v
	}
	return sum / len(numbers)
}

func checkAge(age int) {
	if age < 0 || age > 150 {
        fmt.Println("Enter a valid age range: 0 - 150")
        return
    }
    if age < 18{
        fmt.Println("minor")
    }else if age <=25 { //if(age <= 25){...}
		fmt.Println("2. young adult")
    }else{
        fmt.Println("adult")
    }
}

func reverseString(str string) string{
	var rev string = ""
    for _, v := range(str){
		rev = string(v) + rev
	}
	return rev
}

func findLargestNumber(numbers []int) int{
    large := 0
	for _, v := range(numbers){
		if large < v{
			large = v
		}
	}
	return large
}

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("1. the average number in",slice,"is: ",calculateAverage(slice))
	checkAge(25)
	fmt.Println("3. the reverse of Hello World! : ", reverseString("Hello World!"))
	fmt.Println("4. the largest number in",slice,"is: ",findLargestNumber(slice))
	
	fmt.Println("5. COUNTER")
	counter := Counter{0}
	fmt.Println("COUNTER ADD * 3")
	counter = counter.add()
	counter = counter.add()
	counter = counter.add()
	counter.display()
	fmt.Println("COUNTER SUBTRACT")
	counter = counter.subtract()
	counter.display()
	fmt.Println("COUNTER RESET")
	counter = counter.reset()
	counter.display()
}
