package common

import "fmt"

// NumberDemo This package is for hacking data-types in go-lang
// exploring integer, float, string and properties associated with each type in golang
func NumberDemo() {
	fmt.Println("### Go Language numeric types ###")
	// unsigned int
	// signed int
	// float-pointing
	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println("Integer overflow demo")
	fmt.Println(intNum)
	fmt.Println("Float demo")
	var floatNum1 float32 = 12345678.9
	floatNum2 := 512.56
	fmt.Println(floatNum1)
	fmt.Println(floatNum2)
	fmt.Println("Adding floating-pointer number with integer and returning float")
	fmt.Printf("float: %v + integer: %v = float: %v", floatNum2, intNum, floatNum2+float64(intNum))
}
