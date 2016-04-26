package main

import "fmt"

// FOOMAKERSTART OMIT
type FooMaker interface {
	MakeFoo(input string) string
}

type BarMaker interface {
	MakeBar(input string) string
}

type MyFooBarMaker struct {
}

func (m *MyFooBarMaker) MakeFoo(input string) string {
	return "foo" + input
}

func (m *MyFooBarMaker) MakeBar(input string) string {
	return "bar" + input
}

// FOOMAKEREND OMIT

// FOOMAKERPLAYSTART OMIT
func FooBarSalad(f FooMaker, b BarMaker) string {
	return f.MakeFoo(b.MakeBar("salad"))
}

func main() {
	m := &MyFooBarMaker{}
	fmt.Println(FooBarSalad(m, m))
}

// FOOMAKERPLAYEND OMIT
