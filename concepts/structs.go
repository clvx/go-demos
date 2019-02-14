package concepts

import (
	"fmt"
)

//Struct type
type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

//immutable
func SuperImmutable(s Saiyan) {
	s.Power += 10000
}

//Mutable
// *Saiyan means a pointer to value of type Saiyan
func SuperMutable(s *Saiyan) {
	s.Power += 10000

	//Verifying vegeta's test	
	s = &Saiyan{"Trunks", 1000, nil}
}

//Struct method - receiver
//Receives a pointer to a value.
//Mutable
func (s *Saiyan) SuperMutable() {
	s.Power += 20000
}

// Constructor or a function that returns an instance of the desired type
// Example: A function which returns a Saiyan pointer
// Instead of definying, initializing and getting a struct value we can wrap it
// in a "construct".
func NewSaiyan(name string, power int, sayian *Saiyan) *Saiyan {
	return &Saiyan {
		Name: name,
		Power: power,
		Father: sayian,
	}
}

type Person struct {
	Name string
}

// Person's method
func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s in the compositor method\n", p.Name)
}

//Composition
//Human has a Person type.
type Human struct {
	*Person				//As we do not specify an explicit field name, we can implicitly 
						//access the fields and functions of the composed type.
	Power int
}

//Human's method
func (p *Human) Introduce() {
	fmt.Printf("Hi, I'm %s in the overloaded method\n", p.Name)
}

//Maps
type Namekian struct {
		Name string
		Friends map[string]*Saiyan		//Declares the Friends map. It will panic because it hasn't been initialized.
}
