package test

import (
	"fmt"
	"github.com/zingson/go-helper/hstruct"
	"testing"
)

type A struct {
	Name   string
	Height int64
}

type B struct {
	Name string
	Age  int64
}

func TestAssign(t *testing.T) {
	a := &A{
		Name:   "A",
		Height: 1,
	}
	b := &B{
		Name: "B",
		Age:  10,
	}

	hstruct.Assign(a, b)

	fmt.Printf("%v", b) //&{A 10}
}
