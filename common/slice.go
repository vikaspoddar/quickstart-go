package common

import "fmt"

func SliceDemo() {
	fmt.Println("### Slice Demo ###")
	fmt.Println("Slice related functions: append, copy, len, cap, make, delete")
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	fmt.Println("Append demo")
	intSlice = append(intSlice, 6)
	fmt.Println("len demo")
	fmt.Println(len(intSlice))
	fmt.Println("cap demo")
	fmt.Println(cap(intSlice))
	fmt.Println(intSlice)
	fmt.Println("Adding a slice to another slice with using ... (spread operator)")
	var intSlice2 = []int{7, 8, 9}
	fmt.Println(intSlice2)
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Println("Copy demo")
	var intSlice3 = []int{10, 11, 12}
	fmt.Println(intSlice3)
	copy(intSlice3, intSlice)
	fmt.Println(intSlice3)
	fmt.Println("make function make(sliceType, length, capacity)")
	var intSlice4 = make([]int, 10, 20)
	fmt.Println(intSlice4)
	var intSlice5 []int
	fmt.Println(intSlice5)
	intSlice5 = append(intSlice5, 13)
	fmt.Println(intSlice5)
	intSlice5 = append(intSlice5, 14)
	fmt.Println(intSlice5)
}
