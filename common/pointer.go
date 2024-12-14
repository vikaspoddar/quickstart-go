package common

import "fmt"

func PointerDemo() {
	p := new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)
	*p = 10
	// function-expression
	s := func() string { return fmt.Sprintf("The value of p is: %v", *p) }
	fmt.Printf("%v\n", s())
	p = &i
	fmt.Printf("%v\n", s())
	i = 20
	fmt.Printf("%v\n", s())
}
