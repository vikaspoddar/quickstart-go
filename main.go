package main

import (
	"fmt"
	"github.com/vikaspoddar/go-in-1hr/common"
)

func main() {
	// main package act entry-point to execution
	// every demo has its own package
	// for example
	// 1. testing package: This package demonstrate how simple testing works in go
	common.TickerDemo()
}

//following function are used for testing github-action's CI workflow

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s\n", name)
}
