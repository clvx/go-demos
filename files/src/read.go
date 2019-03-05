package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"path/filepath"
	"io"
)


//Reading executes a set of reading functions
func Reading(path string) {

	//reading data from a file
	data, err := ioutil.ReadFile(filepath.Join(path, "/data"))	// returns ([]byte, err)
	fmt.Println(err)
	Check(err)
	//converting []byte to string.
	fmt.Println(string(data))
	//binary.Write(os.Stdout, binary.LittleEndian, data)

	f, err := os.Open(filepath.Join(path, "/data"))
	Check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	Check(err)
	//Why does it return b1="some "? 
	//Idea: As a slice is sent as a parameter, and a slice is a reference of an array.
	// Read() might add values to the array.
	fmt.Printf("%d bytes: %s\n", n1, string(b1))		//<- This prints 5 bytes: some

	//Seek in a known location in a file and read from there.
	o3, err := f.Seek(6,0)
	Check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//Rewinding a file pointer
	//In other words, taking a file pointer to the beginning of a file.
	_, err = f.Seek(0, 0)
	Check(err)

	//Reading line by line
	s := bufio.NewScanner(f)
	for s.Scan(){
		fmt.Println(s.Text())	
		Check(s.Err())
	}
	
	//Rewinding a file pointer
	_, err = f.Seek(0, 0)
	Check(err)

	//Implementing a buffered reader.
	//bufio allows to read files in chunks.
	r4 := bufio.NewReader(f)	//Creates a new buffered reader
	b4 := make([]byte, 5)		//Byte slice. len=0, cap=5
	for {
		n4, err := r4.Read(b4)
		if err != nil {
			break
		} 
		fmt.Printf("bufio = %d bytes: %s\n", n4, string(b4))
	}
	f.Close()

}
