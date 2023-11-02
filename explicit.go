/* Explicit variable declaration */
package main

import (
	"fmt"
	"reflect"
)

func explicitBasicDataTypes() {
	// strings

	var myWord string = "Hello"
	fmt.Println(myWord)

	// integers
	var x int = 5
	var y int = 10
	var myNum int = x + y
	fmt.Println(myNum)

	// multi-assignment
	var a, b, c = 12, 11, 19
	fmt.Println(a + b + c)

	// boolean
	var myBool bool = true
	fmt.Println(reflect.TypeOf(myBool))

	// array
	var myArray = [3]string{"Volvo", "BMW", "Ford"}

	// array indexation
	fmt.Println(myArray[1])

	// changing a value of an array
	myArray[0] = "Opel"
	fmt.Println(myArray)

}

func explicitConditions(x, y int) {
	/* if - else if - else conditions */
	if x > y {
		fmt.Println("x is greater")
	} else if x == y {
		fmt.Println("equal values")
	} else if x < y {
		fmt.Println("y is greater")
	} else {
		fmt.Println("Not found")
	}
}

func explicitForLoopIteration() {
	/* iteration over an array */
	var myArray = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}

func explicitForLoopWithoutIterator(n int) {
	/* for loop without an iterable. same as while i < n */

	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func explicitFlowControl() {
	/* break if i = 5, skip if i = 3 */

	for i := 0; i < 10; i++ {

		if i == 5 {
			break
		} else if i == 3 {
			continue
		} else {
			fmt.Println("explicitFlowControl | Result: ", i)
		}
	}
}

func main() {
	const numOne int = 5
	const numTwo int = 3

	explicitBasicDataTypes()
	explicitConditions(numOne, numTwo)
	explicitForLoopIteration()
	explicitForLoopWithoutIterator(numTwo)
	explicitFlowControl()
}
