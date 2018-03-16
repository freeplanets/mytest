package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
