package common

import "fmt"

func InputDemo() {
	var input string
	fmt.Printf("Enter your input: ")
	_, _ = fmt.Scanln(&input)
	fmt.Printf("The input is: %s\n", input)
}
