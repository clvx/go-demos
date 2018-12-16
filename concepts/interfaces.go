package concepts

import (
	"fmt"
)

// Interfaces are types that define a contract but not an implementation.
// Interfaces help decouple your code from specific implementation.
// Interfaces are implemented implicitly. It's not needed to specify type T 
// implements interface I. That work is done by the Go compiler.
type I interface {
	f1() string
}

type J interface {
	f2()
}

// Type T satifies interface I.
// Values of type T can be passed to any function accepting I as a parameter.
type T struct{
	Name string // T, and Name are uppercase to be public.
}

type U struct {
	Last string 
}

func (t T) f1() string{
	return t.Name
}

func (u U) f1() string{
	return u.Last
}

func (t T) f2() {
	fmt.Println("f2")
}

// A and B are uppercase to be public.
func A(i I){
	fmt.Println("Hi, you can call me: ", i.f1())
}

func B(j J){
	j.f2()	
}
