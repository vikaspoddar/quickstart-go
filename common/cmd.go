package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type Person struct {
	name   string
	age    int
	gender string
}

func Runner() {
	p := new(Person)
	p.name = "vikas poddar"
	p.age = 18
	p.gender = "male"
	np := *p
	vals := make([]int, 5)
	fmt.Println(vals)
	fmt.Println(np)
}

func Capitalized(name string) (string, int, error) {
	if name == "" {
		return "", 0, errors.New("no name provided")
	}
	return strings.ToTitle(name), len(name), nil
}

func Foo(input int) string {
	isFoo := (input % 3) == 0
	if isFoo {
		return "foo"
	}
	return strconv.Itoa(input)
}

func TestFoo(t *testing.T) {
	result := Foo(9)
	if result != "foo" {
		t.Errorf("expected foo, got %s", result)
	}
}
