package main

import (
	"fmt"
	"go-in-1hr/common"
)

func main() {
	var name string

	fmt.Printf("Enter your name: ")
	_, _ = fmt.Scanln(&name)
	name, size, err := common.Capitalized(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Capitalized name:", name, "Length:", size)
	fmt.Println("Hello World")
}
