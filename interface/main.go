package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

type StructB struct {
	val int
}

func (s *StructB) String() string{
	return "StructB :" + strconv.Itoa(s.val)
}

func (a *StructA) AAA(x int) int {
	return x*x
}

func (s *StructA) String() string {
	return "Val: " + s.val
}

type Printable interface {
	String() string
}

func Println(p Printable) {
	fmt.Println(p.String())
}

func main() {
	a := &StructA{val:"AAA"}
	Println(a)

	b := &StructB{val:10}
	Println(b)
}
