package main

import (
	"fmt"
)
/**
 - Variables with package level scope use var
 */

var nn = 155

func main() {
	/**
	- To local variables we use marmot := to set value
	 */
	a := 15
	sumIt(1,nn+a)
}

 func sumIt(n1 int, n2 int){
	 fmt.Print(n1 + n2)
}