package common

import "fmt"

func VariadicFunctionDemo(nums ...int) {
	fmt.Printf("type of %v: %T\n", nums, nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("sum of the following %v: %v", nums, total)
}
