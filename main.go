package main

import (
	"fmt"
	//"os"
	"github.com/clvx/go-demos/concepts"
)

func main() {
	/*
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	fmt.Println("it's over", os.Args[1])
	*/

	/*
	<STRUCTS>
	*/
	
	//Defining and assigning a Saiyan struct
	goku := concepts.Saiyan{"Goku", 9000, nil}

	//Immutable. Argument is a copy of the memory address(value).
	// There's no change in the calling object.
	concepts.SuperImmutable(goku)
	fmt.Println("Super: ", goku.Power)

	//Mutable. It will change objects values.
	concepts.SuperMutable(&goku)					//Getting the value memory address
	fmt.Println("Super2: ", goku.Power)

	//Calling a struct method
	goku.SuperMutable() //Using methods
	fmt.Println("Super method: ", goku.Power)

	broli := new(concepts.Saiyan)  //Allocates the memory required by a type.
	broli.Name = "Broli"
	broli.Power =  9001	  
	broli.Father = nil
	/**
	 Same as:
	broli := &Saiyan{
		Name: "Gohan",
		Power: 9001,
	}
	**/

	//Struct field in struct
	//Father is a Sayian type.
	gohan := &concepts.Saiyan{
		Name: "Gohan",
		Power: 9001,
		Father: &goku,	//Passing a memory address because the struct defines a struct type.
	}
	fmt.Println("Gohan father: ", gohan.Father.Name)

	//Composition
	krilin := &concepts.Human{
		Person: &concepts.Person{"Krilin"},
		Power: 100,
	}
	krilin.Introduce() //Calling the Human method.
	krilin.Person.Introduce() //Calling the composition Person method.
	fmt.Printf("Krilin as an implicit Human field name: %s\n", krilin.Name)				// Go gives an implicit field name 
																						//to Human using the composition field
	fmt.Printf("Krilin as an explicit field Person name: %s\n", krilin.Person.Name)		// We can access the composition 
																						//variable directly too
	/*
	</STRUCTS>
	*/

	//Arrays 

	var list [10]int	//Declares a fixed array from 0-n
	list[0] = 339		//Assigns foo to the first index in the array

	other_lists := [4]int{9001, 9333, 212, 33} //Declares([n]type) and assigns values({n1, n2, n..}) 
												//to an array from 0-n
	for index, value := range other_lists{
		fmt.Println("printing values: ", index, value)	
	}

	//Slices
	grades := []int{1,4,293,4,9}	//Declares variable, initializes the slice,  
									//and assigns values.
									//length: 5		//slice size
									//capacity: 5	//underlying array size
	other_grades := make([]int, 10)	//Allocate the memory for the underlying array 
								//and also initializes the slice.
								//length: 10
								//capacity: 10
	fmt.Println(grades, other_grades)

	scores := make([]int, 0, 10)	//length: 0
										//capacity: 10
	scores = append(scores, 5)	//Appends in the first free position 
								//in this case: scores[0]
								//If the underlying arrays is full, 
								//it will create a new larger array and copy the values over.
								//Arrays grow with a 2x algorithm.
	scores = scores[0:8]		//Reslicing scores.
								//lenght: 8
								//capacity: 10
	scores[7] = 9033
	fmt.Println(scores, "lenght: ", len(scores), "capacity: ", cap(scores))
	scores = append(scores, 3, 4, 5) //increasing the array over its capacity
										//length: 11
										//capacity: 20 //Remember the 2x algorithm
	fmt.Println(scores, "lenght: ", len(scores), "capacity: ", cap(scores))

	other_grades = append(other_grades, 100)	//increases the capacity from 10 to 20
												//appends in position 11: other_grades[10]
	fmt.Println(other_grades, "lenght: ", len(scores), "capacity: ", cap(scores))

	//COMMON WAYS TO INITIALIZE A SLICE
	names := []string{"leto", "jessica", "paul"}	// To use when you know the values ahead of time.
	checks := make([]bool, 10)						//To use when you write into specific indexes of a slice.
	var students [] string							//nil slice mostly used in conjnunciton with append when 
													//the number of elements is unknown.
	classrooms := make([]int, 0, 20)				//Useful if we have a general idea of how many elements we need.
	fmt.Println(names, checks, students, classrooms)

	//Maps

	lookup := make(map[string]int)	//Allocates and initializes a hash map data structure 
									//and returns a map value that points to it.
	lookup["goku"] = 9000
	power, exists := lookup["vegeta"]
	fmt.Println("Does vegeta have powers?", power, "\nDoes vegeta exist at all?", exists)
	fmt.Println("How many sayians exist?", len(lookup))
	delete(lookup, "goku")	//deleting a value in a map

	type Namekian struct {
		Name string
		Friends map[string]*concepts.Saiyan
	}
	piccolo := &Namekian{
		Name: "Piccolo",
		Friends: map[string]*concepts.Saiyan{					//initializes the Friends map and
				"gohan": &concepts.Saiyan{"Gohan", 9001, nil,},	//adds values to Friends map
			},
		//Friends: make(map[string]*Saiyan),			//ONLY initializes the Friends map
	}
	//piccolo.Friends["gohan"] = &Saiyan{"Gohan", 9001, nil,} //Assigns values to Friends map

	fmt.Println(piccolo.Name, piccolo.Friends["gohan"].Name, piccolo.Friends["gohan"].Power)	//printing dictionary struct values
	//Listing all values of Friends dictionary
	for key, value := range piccolo.Friends {
		println("key: ", key, "values: ", value.Name, value.Power)	//As the map value is a Sayian struct, print the 
																	//struct variables of each key.
	}

	//INTERFACES

	t := concepts.T{Name: "Michael"}
	u := concepts.U{Last: "Ibarra"}
	// A and B implement one type with multiple interfaces.
	concepts.A(t) //Implements interface I of type T - print "you can call me Michael"
	concepts.B(t) //Implements interface J of type T - prints "f2"
	concepts.A(u) //Implements interface I of type U - prints "you can call me Ibarra" 
	concepts.C(t) //Implements interface K of type T using interface J - prints "f2"

	var i concepts.I = concepts.T{Name: "Kelsey"}	//I interface type value.
	fmt.Println(i.F1())								//It can hold any value implementing I.
													//Interface type value gives access ONLY to methods of its 
													// interface type. It hides all the details about the exact 
													// value like if it's a struct, array, scalar, etc.
	var willchange concepts.I = t
	fmt.Println(willchange)	//printing I interface for T type.
	willchange = u			//Changing type from T to U.
	fmt.Println(willchange) //printing I interface for U type.

	//Testing empty interface
	// For an interface to be empty its dynamic type and value need to be nil.
	var m *concepts.T
	if m == nil {
		fmt.Println("m is nill")	 //m is nill
	} else {
		fmt.Println("m is not nill")	
	}
	var n concepts.I = m
	if n == nil {
		fmt.Println("n is nill")	
	} else {
		fmt.Println("n is not nill")	//Even though dynamic value is nill, n 
	}									//is not nill because its dynamic type is 
										//*concepts.T

	//Empty interface
	var k concepts.K
	concepts.Describe(k)

	//Type assertion
	// A way to retrieve dynamic value from interface type value

	var v1 concepts.L = concepts.T{Name: "Luis"}
	v2 := v1.(concepts.T)							//This statement asserts that the interface v1 holds
													//the concrete type T and assigns the underlying 
													//T value to the variable v2.
	//v2 := v1.(concepts.U)							//THIS WILL PANIC
													//If v1 does not hold a T, the statement will trigger a panic.

	//var v1 I										//If v1 is nill then type assertion ALWAYS FAILS.
	//v2 := v1.(concepts.T)							//THIS WILL PANIC	

	fmt.Printf("%T, %s\n", v2, v2.Name)				//prints: (<nil>, <nil>)concepts.T, Luis

	v3, ok := v1.(concepts.T)						//Type assertion can return two values:
													// - The underlying value.
													// - A boolean value that reports wheter the assertion succeded.
													//If v1 holds a T, then v1 will be the underlying value 
													//and ok will be true.
													// If not, ok will be false and v1 will be the zero value of 
													// type T(nill in this case), and no panic occurs.
	//v3, ok := v1.(concepts.U)						//THIS WILL RETURN v2=nill, ok=false
													
	fmt.Printf("%T, %b, %s\n", v3, ok, v3.Name)		//prints: (<nil>, <nil>)concepts.T, %!b(bool=true), Luis

	//Type switch
	//It's like a regular switch statement, but the cases in a type switch specify types(not values), 
	// and those values are compared against the type of the value held by the given interface value.

	var v4 concepts.I
	concepts.Find_type(v4)					//prints nil because the interface type value of v4 is nil.

	var v5 concepts.I = concepts.T{"Duke"}
	concepts.Find_type(v5)					//prints T because the interface type value of v5 is T.


	//Errors
	concepts.Convert("100")					//Converts 100 string to 100 int.
	concepts.Convert("abc")					//Fails gracefully converting 100 string to 100 int.
	concepts.Convert("8")					//Fails gracefully with "Number too small"
	concepts.Convert("11")					//Converts 11 string to 11 int successfully.
	concepts.Convert("51")					//Fails gracefully with "Invalid Range"

	//Defer
	/**
	Immediately after getting a file object with
    `createFile`, we defer the closing of that file
    with `closeFile`. This will be executed at the end
    of the enclosing function (`main`), after
    `writeFile` has finished.
	**/
    f := concepts.CreateFile("/tmp/defer.txt")
    defer concepts.CloseFile(f)
    concepts.WriteFile(f)	

	//Basics
	concepts.Basics("theory")
	concepts.Basics("functions")
	
}
