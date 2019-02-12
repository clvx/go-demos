package concepts

import (
	"fmt"
)

//Basics prints common concepts
func Basics(subject string){
	theory		:=		"golang is a compiled(translating source code to a lower level language), statically typed language(variables MUST be of a specific type[int, string, bool, etc] with a C-like syntax and garbage collection(to keep track of variables and free them when they're no longer used)"
	main		:=		"In Go, the entry point to a program has to be a function called main within a package main"
	imports		:=		"The import keywork is used to declare the packages that are used to declare the packages that are used by the code in the file. Go is strict about importing packages. It will not compile if you import a package but don't use it."
	declaration :=		"Use var NAME TYPE when declaring a variable to its zero value"
	assignment  :=		"NAME := VALUE when declaring and assigning a value, and NAME = VALUE when assigning to a previously declared variable"
	functions	:=		`
func log(message string) {}				<- No return value
func add(a int, b int) int {}			<- Returning 1 value
func power(name string) (int, bool) {}	<- Returning 2 values
value, exists := power("foo")			<- Using the power func which returns two values.
_, exists := power("bar")				<- _, blank identifier, the return value is not assigned.
`
	structs		:=		`
A type which contains named fields. Fields can be of any type.
type X struct{
	foobar String,
}
type Name struct{
	foo int,
	bar *X,
}

name := &Name{1, &bar{"Hola mundo"}}
name := &Name{foo: 1, bar: &bar{"hello world"}}
`
	pointers	:=		`
A pointer is a memory address. It's the location of where to find the actual value. It's a level of indirection. 
type Name struct{
	foo int,
}
foobar := Name{1}				<- Assigning a value to foobar
func adding(n Name){
	n.bar += 1
}
foobar := &Name{1}				<- Getting and assigning the address of our value(& address of operator) to foobar. References needs to be an address type of *Name where *X means pointer to a value of X.
func adding(n *Name){
	n.bar += 1
}

Function arguments are passed as copies. Copying a value can be inefficient. To avoid this we use pointers to reference a value instead of copyting values.
However, values are greate to make data immutable(changes that a function makes to it won't be reflected in the calling code).
`
	methods		:=		`
Name is the receiver of the Adding method
type Name struct {
	bar int,
}

func (n *Name) Adding() {
	n.bar += 1
}
foo := &Name{1}
foo.Adding()
`
	constructors	:=		`
Structs don't have constructors; instead, you create a function that returns the desired type. It can be a pointer or an address.
type NewName(number int) *Name {
	return &Name{
		foo int,	
	}
}

Using new() built-in function allocates the memory required by a type instead of making your own initialization functions.

`
	composition		:=  `
	
Golang does not support inheritance; however, it does support for composition combining simple types to make more complex ones. This type HAS another type.
type Car struct {
	model string,
}

type Wheel struct {
	quantity	int,
	*Car,			<- Composition right here. A car has wheels.
}

func (c *Car) Show () {
	fmt.Printf("It's a %s", c.model)
}

v := &Wheel{10, &Car{sedan}}
v.Show()			<- It should print "It's a sedan"
`

	overloading		:=	`
Go does not support overloading, but we can overwrite the functions of a composed type.

//Review concept composition for full example
func (w *Wheel) Show () {
	fmt.Printf("It has %d wheels", w.quantity)
}

v.Show()			<- It should print "Wheel quantity" 
v.Car.Show()		<- It should print "It's a sedan"

`

	arrays			:= `
	Declaring an array requires that we specify the size, and once the size is specified, it cannot grow.
	var NAME [n]type
	NAME	:= [n]type{v1, v2, ..., vn}
	NAME[i] = SOME_VALUE
	
`

	switch s := subject; s {
		case "theory":
			fmt.Println(theory)
		case "main":
			fmt.Println(main)
		case "import":
			fmt.Println(imports)
		case "declaration":
			fmt.Println(declaration)
		case "assignment":
			fmt.Println(assignment)
		case "functions":
			fmt.Println(functions)
		case "structs":
			fmt.Println(structs)
		case "pointer":
			fmt.Println(pointers)
		case "methods":
			fmt.Println(methods)
		case "constructors":
			fmt.Println(constructors)
		case "composition":
			fmt.Println(composition)
		case "overloading":
			fmt.Println(overloading)
		case "arrays":
			fmt.Println(overloading)
		default:
			fmt.Println("Concept not found")
	}
}
