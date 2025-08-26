package main

import (
	math "Day_1/Day_1/Math"
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Enter the value of a: ")
	fmt.Scanln(&a)

	fmt.Print("Enter the value of b: ")
	fmt.Scanln(&b)

	fmt.Printf("Sum of (a, b) = %d\n", math.Add(a, b))
	fmt.Printf("Multiply of (a, b) = %d\n", math.Multiply(a, b))

}
