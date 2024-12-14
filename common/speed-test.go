package common

import (
	"fmt"
	"time"
)

func SpeedTest() {
	fmt.Println("### Speed test ###")
	size := 1000_000
	var testSlice1 []int
	var testSlice2 = make([]int, 0, size)
	fmt.Println("Time to append to slice 1 (without pre-allocation):", timeLoop(testSlice1, size))
	fmt.Println("Time to append to slice 2 (with pre-allocation):", timeLoop(testSlice2, size))
}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
