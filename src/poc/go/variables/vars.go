package main

import (
	"fmt"
	_ "fmt"
)

var day = "friday"

func main() {
	name := "Bryan"
	age := 29
	wantCoffee := true
	/**
	*** use of different types of variables
	 */
	fmt.Print("name: ", name, "\n Age: ", age, "\n Want to drink coffee? ", wantCoffee, "\n")
	/**
	*** return value of variable and type
	 */
	fmt.Printf("age= %v, %T", age, age)
	fmt.Print("\n")

	/**
	*** return data and errors
	 */
	data, _ := fmt.Print(age)
	fmt.Print(data, "\n")

	/**
	*** multiple declarations
	 */

	x, y := 202, 10
	fmt.Print(sum(x, y), day)

}

func sum(a int, b int) int {
	return a + b
}
