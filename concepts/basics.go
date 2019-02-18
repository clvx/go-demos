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

	slices			:= `
	A slice is a lightweight structure that wraps and represents a portion of an array.
	length: slice size. The lenght cannot be greater than the capacity.
	capacity: underlying array size.
	NAME := []type{v1, v2,v3, ..., vn}	  //Allocates memory for the underlying array.
										  //length: n
										  //capacity: n
	var NAME []type						  //Allocates memory for the underlying array.
										  //length: 0
										  //capacity: 0
	NAME := make([]type, n, m)			  //Allocate the memory for the underlying array 
										  //and also initializes the slice.
										  //length: n
										  //capacity: m - optional

	Adding elemets to a slice:
	vector := make([]type, n, m)
	vector = append(vector, some_value)	  //If the underlying array is full, it will create a new larger array and copy the values over.
										  //Arrays grow with a 2x algorithm. E.g. if capacity initially is 5, then it will be 10, then 20 and so on.
	- If lenght is lenght(n) is 0, it will add the new element at index 0. To add an element in any other index, re-slicing the slice is necessary:
	vector = vector [0:i]			      //initializing i elements with zero value(0, 0, ..., i). 
	vector[i] = another_value			  //slices don't create a new array with copies of the values; instead, it works on the array values.
	vector[:i]	//shorthand for from the start to i.
	vector[i:]  //shorthand for from i to the end.
									
`
	maps			:=`
	A map is an unordered collection of key-value pairs, where each key is unique.
	var NAME map[type]type						//Declares an empty map with zero values(nill). A panic will occur if items are added without the map being initialized.
												//length: 0
												//capacity: 0
	NAME := make(map[type]type, n)				//Declares and initializes a nill map
												//length: 0
												//capacity: n. Optional, in case of absence the capacity is 0.
	NAME := map[type]type{}						//Declares and initializes a nill map
												//length: 0
												//capacity: 0
	NAME := map[type]type{k1:v1, .., kn:vn}		//Declares, initializes, and assigns values to a map
												//length: n
												//capacity: n
	NAME[key] = value							//Assigning a value to a key
	delete(NAME, key)							//Deletes a key

	for key, value := range NAME {
		fmt.Println(key, value)					//If a value doesn't exist the zero value of the type is returned.
	}

	value, ok := NAME[key]						//Returns zero or value and a boolean(true or false).

	A map is a reference type. When a map is assigned to a new variable, both refer to the same underlying data structure. Therefore changes done by one variable will be visible to the other.
`
	packages		:= `
	Packages follow the directory structure of your GO workspace. The name of the package is the same as the name of the folder. Types and functions are visibiles as long as they start with an uppercase Letter.
	$cat utils/foo.go
	package utils

	//Public variable
	type Foo struct{
		name string,	
	}

	//Private variable
	type fooBar struct{
		name string,	
	}

	$cat main.go
	import (
		"utils"
	)

	some_var := &utils.Foo{"bar"}

	The name of the package is the same as the name of the folder.
	
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
			fmt.Println(arrays)
		case "slices":
			fmt.Println(slices)
		case "maps":
			fmt.Println(maps)
		case "packages":
			fmt.Println(packages)
		default:
			fmt.Println("Concept not found")
	}
}
