package common

import "fmt"

func ArraysDemo() {
	fmt.Println("### Arrays Demo ###")
	fmt.Println("Arrays are fixed length collection of data of same type")
	fmt.Println("Arrays are zero-indexed")
	fmt.Println("Arrays are mutable")
	fmt.Println("Arrays are not null terminated")
	array := [5]string{"fixed length", "same type", "indexable (0-indexed)", "contiguous memory", "mutable"}
	for index, value := range array {
		fmt.Printf("array[%v] = %v\n", index, value)
	}
	fmt.Println("Array memory address")
	for index, _ := range array {
		fmt.Println(&array[index])
	}
	fmt.Printf("Arrays [1:3] %v\n", array[1:3])
	fmt.Println(" [...]int{value1, value2, ...} is a shorthand for [n]int{value1, value2, ...}")
}
