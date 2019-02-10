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
		default:
			fmt.Println("Concept not found")
	}
}
