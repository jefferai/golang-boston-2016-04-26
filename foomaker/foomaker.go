package main

import "fmt"

// FOOMAKERPLAYSTART OMIT
// FOOMAKERSTART OMIT
type FooMaker interface {
	MakeFoo(input string) string
}

type MyFooMaker struct {
}

func (m *MyFooMaker) MakeFoo(input string) string {
	return "foo" + input
}

// FOOMAKEREND OMIT
func FooSalad(f FooMaker) string {
	return f.MakeFoo("salad")
}

func main() {
	m := &MyFooMaker{}
	fmt.Println(FooSalad(m))
}

// FOOMAKERPLAYEND OMIT
