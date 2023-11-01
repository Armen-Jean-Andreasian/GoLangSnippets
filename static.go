package main

import (
	"fmt"
)

func staticBasicDataTypes() {
	var myNum int = 23
	var myWord string = "Hello"

	fmt.Println(myNum)
	fmt.Println(myWord)

	var x int = 5
	var y int = 10
	var z int = x + y

	fmt.Println(z)

	var myBool bool = true
	fmt.Println(myBool)
}

func myStaticArrays() {
	var myArray = [3]string{"Volvo", "BMW", "Ford"}
	fmt.Println(myArray)
}

func main() {
	staticBasicDataTypes()
	myStaticArrays()
}
