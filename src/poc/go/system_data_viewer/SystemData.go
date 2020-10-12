package main

import (
	"fmt"
	_ "fmt"
	"runtime"
)

func main(){
	fmt.Println("Operating System")
	fmt.Println(runtime.GOOS)
	fmt.Println("Architecture")
	fmt.Println(runtime.GOARCH)
}