package src

import (
	"bufio"
	"fmt"
	"errors"
	"strings"
)

/**
	Package bufio helps with buffered I/O. 
	Buffered I/O refers to the technique of temporarily storing the results of an 
	I/O operation in user-space before transmitting it to the kernel (in the case of writes) 
	or before providing it to your process (in the case of reads). By so buffering the data, 
	you can minimize the number of system calls and can block-align I/O operations(in the 
	case of block devices), which may improve the performance of your application.
	Source: https://www.quora.com/In-C-what-does-buffering-I-O-or-buffered-I-O-mean/answer/Robert-Love-1

	bufio.Writer: buffering 4 bytes
	producer	---->	buffer	---->	consumer/destination(io.Writer)
	
	a			---->	a		---->	
	b			---->	ab		---->	
	c			---->	abc		---->	
	d			---->	abcd	---->	
	e			---->	e		---->	abcd
	f			---->	ef		---->	abcd
	g			---->	efg		---->	abcd
	h			---->	efgh	---->	abcd
	i			---->	i		---->	abcdefgh
**/


//E1 simulation bufio.Writer
//Writer from bufio.Writer
//type Writer struct {
//	err error			//After encountering an error, writer is no-op
//	buf []byte			//Accumulates data
//	n int
//	wr io.Writer		//Gets data when buffer is full or Flush() is called
//}
type E1 int

func (*E1) Write(p []byte) (n int, err error) {
	fmt.Printf("Write # of bytes: %d, value: %q \n", len(p), p) 
	return len(p), nil
}

//E2 implements Writer
type E2 int

func (*E2) Write(p []byte) (n int, err error) {
	fmt.Printf("Write: %q\n", p) 
	return 0, errors.New("Boom. following Flush() does not execute due error")
}

//Execute executes the bufio.Writer example
func Execute(){
	//Each write operation goes straight to destrination.
	fmt.Println("Unbuffered I/O")
	w := new(E1)
	w.Write([]byte{'a'})	//writes
	w.Write([]byte{'b'})	//writes
	w.Write([]byte{'c'})	//writes
	w.Write([]byte{'d'})	//writes
	//Each write operation goes to the buffer(size 3), when buffer is full it 
	//goes to destination.
	fmt.Println("Buffered I/O")
	//returns a new Writer pointer whose buffer has at least the specified size.
	buf := make([]byte, 0)
	bw := bufio.NewWriterSize(w, 3)
	for i, s := range "efgh" {
		fmt.Printf("buffers taken: %d, buffers available: %d\n, ", bw.Buffered(), bw.Available())	//it starts at 0
		bw.Write(append(buf, byte(s)))	
		buf = buf[:0]
		if i >= 3 {
			fmt.Println("Consumer was called. Buffer is empty now")
			fmt.Printf("writes and buffers: %d\n", bw.Buffered())	
		}
	}
	//Flush writes any buffered data to the underlying io.Writer.
	//Flush is a receiver expecting a *Writer
	err := bw.Flush()
	Check(err)


	/**
	Unbuffered I/O
	Write # of bytes: 1, value: "a" 
	Write # of bytes: 1, value: "b" 
	Write # of bytes: 1, value: "c" 
	Write # of bytes: 1, value: "d" 
	Buffered I/O
	buffers: 0
	buffers: 1
	buffers: 2
	buffers: 3
	Write # of bytes: 3, value: "efg" 
	Consumer was called. Buffer is empty now
	writes and buffers: 1
	Write # of bytes: 1, value: "h" 
	**/

	w2 := new(E2)
	//Reset discards any unflushed buffered data, clears any error, and resets 
	//current writer to be used with a different writer consumer.
	bw.Reset(w2)
	s2 := strings.NewReader("abc") 
	//ReadFrom:
	//- reads data from a reader when buffer is filled it call the consumer/writer
	//- implements io.ReaderFrom
	bw.ReadFrom(s2)
	bw.Write([]byte{'d'})	//writes and buffers
	err = bw.Flush()
	fmt.Println(err)

	/**
	Write: "abc"
	Boom. following Flush() does not execute due error"
	**/
}
