package concepts

import (
	"fmt"
	"strconv"
	"errors"
)

//It creates a new struct for errors.
type ErrInvalid struct {
	message string
}

// ErrInvalid implements the Error interface.
func (e *ErrInvalid) Error() string {
	return e.message
}

// Implementing a function to set up the ErrInvalid.message value.
// This function can be avoided if the message value is declared directly in the 
// ErrInvalid struct.
func NewErrInvalid(message string) *ErrInvalid {
	return &ErrInvalid{
		message: message,	
	}
}

var ErrInvalidRange = errors.New("Invalid Rage") //Defining an error with the errors package.

func Convert(number string) {
	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("not a valid number")
	} else if n <= 10 {
		fmt.Println(NewErrInvalid("Number too small"))
		//fmt.Println(&ErrInvalid{"Number too small"})	//This implements the interface without the need of NewErrInvalid() 
	} else if n > 50 && n <= 60 {
		fmt.Println(ErrInvalidRange)
	} else {
		fmt.Println(n)	
	}
}
