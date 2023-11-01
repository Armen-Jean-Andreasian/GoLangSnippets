package main

import "fmt"

func dynamicBasicDataTypes() {
	// comment
	var myNum = 23
	var myName = "Hello world"

	fmt.Println(myNum)
	fmt.Println(myName)

	var x = 5
	var y = 10
	var z = x + y

	fmt.Println(z)

	var a, b, c = 12, 11, 19
	fmt.Println(a + b + c)
}

func myDynamicArrays() {
	var myArray = []string{"Volvo", "BMW", "Ford"}
	fmt.Println(myArray)
}

func main() {
	dynamicBasicDataTypes()
	myDynamicArrays()
}
