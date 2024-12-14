package common

import (
	"fmt"
	"time"
)

func TimeDemo() {
	fmt.Println("### Time Demo ###")
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("It's Saturday")
	default:
		fmt.Println("It's not Saturday")
	}
}
