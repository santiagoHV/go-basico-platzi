package main

import "fmt"

func normalFunction() {
	fmt.Println("Hello, World!")
}

func functionWithParameters(a, b int) { // a and b are integers
	fmt.Println(a + b)
}

func functionWithReturnValue() int {
	return 5
}

func functionWithMultipleReturnValues() (int, int) {
	return 5, 6
}

func main() {
	fmt.Println("Hello, World!")

	normalFunction()
	var value int = functionWithReturnValue()
	fmt.Println(value)

	var value1, value2 int = functionWithMultipleReturnValues()
	fmt.Println(value1, value2)
	value3, _ := functionWithMultipleReturnValues() // _ is used to ignore the return value
	fmt.Println(value3)

}
