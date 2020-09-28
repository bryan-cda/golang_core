package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	sizeOf, err := fmt.Print("こんいちは - hello \n")
	fmt.Print(sizeOf, err)
}
