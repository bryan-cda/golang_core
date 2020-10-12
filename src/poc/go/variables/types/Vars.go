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
	fmt.Println("")
	x, y := 202, 10
	fmt.Print(sum(x, y), day)

	/**
	*** Print UT8 byte code for characters b r y a n
	 */
	fmt.Println("")
	fmt.Println("Access UTF-8 code of characters of name")
	ut8 := []byte(name)
	fmt.Println(ut8)


}

func sum(a int, b int) int {
	return a + b
}
