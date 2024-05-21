package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator function to perform operations based on the operand
// Calculate function takes 3 arguments: 2 float and a string
// returns the result if calculation is successful else print error
func Calculator(value1, value2 float64, operand string) (float64, error) {
	//switch to identify the operands
	switch operand {
	case "+":
		return (value1 + value2), nil
	case "-":
		return (value1 - value2), nil
	case "*":
		return (value1 * value2), nil
	case "/":
		if value2 == 0 {
			return 0, errors.New("Division by 0 is not possible")
		} else {
			return (value1 / value2), nil
		}
	default:
		return 0, errors.New("Calculation Error: check the operand")
	}
}

func main() {

	//Reading the Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	//Check if exact 3 parts are given. If not, throw an error
	if len(parts) != 3 {
		fmt.Println(errors.New("Non valid Inputs: the input should be exactly 3 parts"))
		return
	}

	//converting intput str to float
	value1, convErr1 := strconv.ParseFloat(parts[0], 64)
	value2, convErr2 := strconv.ParseFloat(parts[2], 64)

	//Handle the error occured while conversion if inputs -- If the error is present
	if convErr1 != nil || convErr2 != nil {
		fmt.Println(errors.New("Unable to convert the string to float -- check the input"))
		return
	}

	result, calculateErr := Calculator(value1, value2, parts[1])
	//check for error if returned by the Calculator function
	if calculateErr != nil {
		fmt.Println(calculateErr)
		return
	}

	//Output the result
	fmt.Printf("Result: %v\n", result)
}
