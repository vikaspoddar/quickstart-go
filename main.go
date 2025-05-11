package main

import (
	"fmt"
	"os"
)

func main() {
	// main package id entry-point to execution
	// every demo has its own package
	// for example
	// 1. testing package: This package demonstrate how simple testing works in go
	name := os.Getenv("NAME")
	greeting := sayHello(name)
	fmt.Println(greeting)
}

//following function are used for testing github-action's CI workflow

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s\n", name)
}
