package concepts

import (
	"fmt"
)

/**
- Interfaces are types that define a contract but not an implementation.
- Interfaces help decouple your code from specific implementation.
- Interfaces are implemented implicitly. It's not needed to specify type T 
implements interface I. That work is done by the Go compiler.
- Declaration of the interface type specifies methods belonging to it. 
- Method is defined by its name and signature. 
- It's allowed to include other interfaces.
- Interface methods need to have UNIQUE names.
- If a method is being added to one of the types implemeting the interface, 
then all types need to implement that method. See F1(), f1(), type T and type U.
- Variables have type known at compilation phase. It’s specified while declaration, 
never changes and is known as static type (or just type). Variables of interface 
type also have static type which is an interface itself. They additionally have 
dynamic type so the type of assigned value.
**/
type I interface {
	f1() string
	F1() string
}

type J interface {
	f2()
}

// All J methods will be available to interface K.
type K interface {
	J
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

func (t T) F1() string{
	return t.Name
}

func (u U) f1() string{
	return u.Last
}

func (u U) F1() string{ //This will error out if missing, 
	return u.Last		//even though only t.F1() is being used.
}

func (t T) f2(){
	fmt.Println("f2")
}

// A and B are uppercase to be public.
func A(i I){
	fmt.Println("Hi, you can call me: ", i.f1())
}

func B(j J){
	j.f2()	
}

func C(k K){
	fmt.Println("Inside K interface, printing J interface f2 function:")
	k.f2()
}
