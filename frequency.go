package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

func getInput(cmd string) string{
	fmt.Printf("%s : ", cmd)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func countFreq(words []string)  map[string]int{
	frequency := map[string]int{}
	for _, v := range(words){
		frequency[v] += 1
	}
	return frequency
}

func displayFreq(frequency map[string]int){
	for i, v := range(frequency){
		fmt.Printf("%s : %d\n", string(i), v)
	}
}

func main(){
	sentence := getInput("Enter the Sentence")
	words := strings.Split(sentence, " ")
	frequency := countFreq(words)
	displayFreq(frequency)
}