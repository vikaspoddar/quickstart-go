package common

import "fmt"

func StringDemo() {
	fmt.Println("### String demo ###")
	fmt.Println("### Important note on string in golang ###")
	fmt.Println("String are not collection of characters, but Utf-8 encoded bytes")
	fmt.Println("String are immutable")
	fmt.Println("String are not null terminated")
	fmt.Println("For learning more about string in golang, its internal implementation, refer to rune type and rune package")
}
