package main

import "fmt"

// FOOMAKERSTART OMIT
type FooMaker interface {
	MakeFoo(input string) string
}

// FOOMAKERPLAYSTART OMIT
type MyBarMaker struct {
	FooMaker
}

func (b *MyBarMaker) MakeBar(input string) string {
	return "bar" + input
}

type MyFooMaker struct {
}

func FooBarSalad(f FooMaker, b *MyBarMaker) string {
	return f.MakeFoo(b.MakeBar("salad"))
}

func main() {
	m := &MyBarMaker{FooMaker: &MyFooMaker{}}
	fmt.Println(FooBarSalad(m, m))
}

// FOOMAKEREND OMIT
func (m *MyFooMaker) MakeFoo(input string) string {
	return "foo" + input
}

// FOOMAKERPLAYEND OMIT
