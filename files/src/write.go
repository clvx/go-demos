package src

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)


//Writing writes to a file
func Writing(path string){

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(filepath.Join(path, "/wdata-1"), d1, 0644)
	Check(err)

	//os.Create() returns a file descriptor(*os.File, int)
	f2, err := os.Create(filepath.Join(path, "/wdata-2"))
	Check(err)

	defer f2.Close()

	//Bytes 0-255. Not like 333 dumbass.
	d2 := []byte{115, 111, 101, 255}
	n2, err := f2.Write(d2)
	Check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//WriteString writes STRIIIIIINGS!!
	//WriteString returns number of bytes written
	n3, err := f2.WriteString("\nsome nonsense string\n")
	fmt.Printf("wrote %d bytes\n", n3)
	f2.Sync()

	//Implementing a buffered writer 
	w := bufio.NewWriter(f2)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush() 

	//Inserting line by line
	f3, err := os.Create(filepath.Join(path, "/wdata-3"))
	Check(err)
	defer f3.Close()
	d := []string{"Welcome to the world of Go", "Go is a compiled language", "It is easy to learn Go."}
	for _, v := range d {
		//fmt.Fprintln takes a io.writer as a parameter and appends a new line.
		_, err = fmt.Fprintln(f3, v)
		Check(err)
	}

	//Appending to a file
	f4, err := os.OpenFile(filepath.Join(path, "/wdata-1"), os.O_APPEND|os.O_WRONLY, 0644)
	Check(err)
	defer f4.Close()
	newLIne := "File handling is easy"
	_, err = fmt.Fprintln(f4, newLIne)
	Check(err)
	fmt.Println("File append successfully")
}
