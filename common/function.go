package common

import (
	"errors"
	"fmt"
)

func FunctionDemo() {
	printMe("Print me Daddy")
	quotient, remainder, err := intDivision(9, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Quotient: %v, Remainder: %v\n", quotient, remainder)
		switch remainder {
		case 0:
			fmt.Println("Perfectly divisible")
		case 1, 2:
			fmt.Println("Almost divisible")
		default:
			fmt.Println("Not divisible")
		}
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
