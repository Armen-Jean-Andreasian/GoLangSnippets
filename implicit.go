/* Implicit variable declaration */
package main

import (
	"fmt"
	"reflect"
)

func implicitBasicDataTypes() {
	// strings

	myWord := "Hello world"
	fmt.Println(myWord)

	// integers
	x := 5
	y := 10
	myNum := x + y
	fmt.Println(myNum)

	// multi-assignment
	a, b, c := 12, 11, 19
	fmt.Println(a + b + c)

	// boolean
	myBool := true
	fmt.Println(reflect.TypeOf(myBool))

	// array
	myArray := []string{"Volvo", "BMW", "Ford"}

	// array indexation
	fmt.Println(myArray[1])

	// changing a value of an array
	myArray[0] = "Opel"
	fmt.Println(myArray)

}

func implicitConditions(x, y int) {
	/* switch - case - default conditions */
	switch {
	case x > y:
		fmt.Println("x is greater")
	case x == y:
		fmt.Println("equal values")
	case x < y:
		fmt.Println("y is greater")
	default:
		fmt.Println("Not found")
	}
}

func implicitForLoopIteration() {
	/* iteration over an array */
	myArray := []int{1, 2, 3, 4, 5}

	for _, value := range myArray {
		fmt.Println("implicitForLoopIteration | Result: ", value)
	}
}

func implicitForLoopWithoutIterator(n int) {
	/* for loop without an iterable. same as while i < n */

	i := 0
	for i < n {
		fmt.Println("implicitForLoopWithoutIterator | Result: ", i)
		i++
	}
}

func implicitFlowControl() {
	/* break if i = 5, skip if i = 3 */

	i := 0
	for i < 10 {
		i++

		if i == 5 {
			break
		} else if i == 3 {
			continue
		} else {
			fmt.Println("implicitFlowControl | Result: ", i)
		}
	}
}

func main() {
	numOne := 5
	numTwo := 3

	implicitBasicDataTypes()
	implicitConditions(numOne, numTwo)
	implicitForLoopIteration()
	implicitForLoopWithoutIterator(numTwo)
	implicitFlowControl()
}
