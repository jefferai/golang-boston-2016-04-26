package main

import "fmt"

// FOOMAKERSTART OMIT
type FooMaker interface {
	MakeFoo(input string) string
}

type BarMaker interface {
	MakeBar(input string) string
}

type MyFooMaker struct {
}

func (m *MyFooMaker) MakeFoo(input string) string {
	return "foo" + input
}

// FOOMAKERPLAYSTART OMIT
type MyFooBarMaker struct {
	MyFooMaker
}

func (m *MyFooBarMaker) MakeBar(input string) string {
	return "bar" + input
}

// FOOMAKEREND OMIT
func FooBarSalad(f FooMaker, b BarMaker) string {
	return f.MakeFoo(b.MakeBar("salad"))
}

func main() {
	m := &MyFooBarMaker{}
	fmt.Println(FooBarSalad(m, m))
}

// FOOMAKERPLAYEND OMIT
