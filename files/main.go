package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
	//"encoding/binary"
	"io"
)

func check(e error){
	if e !=  nil {
		os.Exit(1)			
	}
}

func reading() {
	//reading data from a file
	data, err := ioutil.ReadFile("./data")	// returns ([]byte, err)
	check(err)
	//converting []byte to string.
	fmt.Println(string(data))
	//binary.Write(os.Stdout, binary.LittleEndian, data)

	f, err := os.Open("./data")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	//Why does it return b1="some "? 
	//Idea: As a slice is sent as a parameter, and a slice is a reference of an array.
	// Read() might add values to the array.
	fmt.Printf("%d bytes: %s\n", n1, string(b1))		//<- This prints 5 bytes: some

	//Seek in a known location in a file and read from there.
	o3, err := f.Seek(6,0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//Rewinding a file pointer
	//In other words, taking a file pointer to the beginning of a file.
	_, err = f.Seek(0, 0)
	check(err)

	//Reading line by line
	s := bufio.NewScanner(f)
	for s.Scan(){
		fmt.Println(s.Text())	
		check(s.Err())
	}
	
	//Rewinding a file pointer
	_, err = f.Seek(0, 0)
	check(err)

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

	
func writing(){
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./wdata-1", d1, 0644)
	check(err)

	f2, err := os.Create("./wdata-2")
	check(err)

	defer f2.Close()

	//Bytes 0-255. Not like 333 dumbass.
	d2 := []byte{115, 111, 101, 255}
	n2, err := f2.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//WriteString writes STRIIIIIINGS!!
	n3, err := f2.WriteString("\nsome nonsense string\n")
	fmt.Printf("wrote %d bytes\n", n3)
	f2.Sync()

	w := bufio.NewWriter(f2)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush() 
}

func main() {
	reading()
	//writing()
}
