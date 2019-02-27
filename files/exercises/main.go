package main
/**

Difference between bufio, io, io/ioutil
fmt.Println()
ioutil.Readfile()
io.Readfull()
ioutil.ReadAll()
bufio.NewScanner()
- bufio.Scan()
- bufio.Text()
- bufio.Err()
bufio.NewReader()
- bufio.Readline()

**/

import (
	"bufio"		//Package to read/write buffering
	"fmt"
	"io/ioutil" //Package to read files
	"os"
	"io"
	"strings"
	"math/rand"
	"time"
)

func check(e error) {
	if e != nil {
		os.Exit(1)
	}
}

//ETA 30h

func remove(source []byte, index int) []byte {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func main() {

	//write a program to read an entire text file.
	data, err := ioutil.ReadFile("./foo")
	check(err)
	fmt.Printf("\n1: %s", string(data))

	//Write a program to read first n lines of a file.
	//Buffered reading
	f1, err := os.Open("./foo")
	check(err)
	defer f1.Close()

	n := 19
	bp := bufio.NewReader(f1)
	buffer := make([]byte, n)
	br, err := bp.Read(buffer) //Returns int, but it modifies the underliying array|buffer.
	check(err)
	fmt.Printf("\n2: Read %d bytes: %s", br, string(buffer))

	//Write a program to append text to a file and display the text.
	f3, err := os.OpenFile("./foo", os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer f3.Close()
	line := "Some data to append"
	br1, err := fmt.Fprintln(f3, line)
	check(err)
	data2, err := ioutil.ReadFile("./foo")
	check(err)
	fmt.Printf("\n3: Read %d bytes: %s", br1, string(data2))

	//Write a program to read last n lines of a file.
	//Write a program to get the file size of a plain file.
	//Write a program to find properties of a file using stat() function.
	//If file the file is massive then creating a buffered reader is the best option
	var n4 int64 = 10
	f4, err := os.Open("./foo")
	check(err)
	defer f4.Close()
	f4a, err := os.Stat("./foo")
	check(err)
	f4.Seek(f4a.Size()-n4, 0)	//Reads from total bytes - n
	//buffer4 := make([]byte, n4)									//Long way to do the same
	//br4, err := io.ReadFull(f4, buffer4) 							//Long way to do the same
	//fmt.Printf("\n4: Read %d: %s									//Long way to do the same", br4, string(buffer4))
	br4, err := ioutil.ReadAll(f4)
	fmt.Printf("\n4: Read: %s", string(br4))

	//Write a program to read a file line by line store it into a variable.
	var s5 string
	f5, err := os.Open("./foo")
	check(err)
	defer f5.Close()

	br5 := bufio.NewScanner(f5)
	for br5.Scan(){
		s5 = s5 + "\n" + br5.Text()
		check(br5.Err())
	}
	fmt.Printf("\n5: Printing var line by line:%s", s5)

	//Write a program to read a file line by line store it into a slice.
	//Write a program to count the number of lines in a text file.
	f6, err := os.Open("./foo")
	check(err)
	defer f6.Close()

	br6 := bufio.NewScanner(f6) 
	var v6 []string 
	counter6 := 0
	for br6.Scan(){
		v6 = append(v6, br6.Text())
		check(br6.Err())
		counter6++
	}
	fmt.Printf("\n6: Total of lines: %d\n", counter6)
	
	//Write a program to find the longest words.
	//- ReadAll, break in pieces, convert to string, sort array, get the biggest array.
	//- Better approach with scan.Words(): get first word, scan each word,  
	//compare with first word, if longer replace variable, print variable.
	f7, err := os.Open("./foo")
	check(err)
	defer f7.Close()

	scan7 := bufio.NewScanner(f7)
	scan7.Split(bufio.ScanWords)
	i := 1
	s7 := ""
	tmp7 := ""
	for scan7.Scan(){
		if i == 1 {
			s7 = scan7.Text()
			check(scan7.Err())
		} else {
			tmp7 = scan7.Text()
			check(scan7.Err())
			if len(s7) <= len(tmp7) {
				s7 = tmp7
			}
		}
		i++
	}
	fmt.Printf("\n7: longest word: %s", s7)
	
	//Write a program to write a list to a file.
	//Creates if doesn't exist; otherwise, appends to bar
	f8, err := os.OpenFile("./bar", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer f8.Close()

	data8 := []string{"Some data to dump", "more data to dump"}
	for _, s8 := range(data8){
		n8, err := fmt.Fprintln(f8, s8)
		check(err)
		fmt.Printf("\n8: Writing %d bytes", n8)
	}

	//Write a program to copy the contents of a file to another file.
	//If file the file is massive then creating a buffered reader is the best option
	fr9, err := os.Open("./foo")
	check(err)
	defer fr9.Close()
	
	fw9, err := os.OpenFile("./foobar", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer fw9.Close()

	_, err = io.Copy(fw9, fr9)
	check(err)

	//Write a program to count the frequency of words in a file.
	//Improved logic: Map login should be in scan.Words() loop.
	//Bad Approach: Read file, create string slice, make a map of string: int, insert values in map.

	f10, err := os.Open("./foo")
	check(err)
	defer f10.Close()
	scan10 := bufio.NewScanner(f10)
	scan10.Split(bufio.ScanWords)
	map10 := make(map[string]int)
	for scan10.Scan(){
		map10[scan10.Text()]++
		check(scan10.Err())
	}
	for k10, v10 := range map10{
		fmt.Printf("\n10: key=%s, value=%d", k10, v10)
	}

	//Write a program to read a random line from a file.
	fr11, err := os.Open("./foo")
	check(err)
	defer fr11.Close()

	s11, err := os.Stat("./foo")
	check(err)

	rand.Seed(time.Now().UnixNano())
	n11 := 0 + rand.Intn(int(s11.Size())-0+1)
	fr11.Seek(int64(n11), 0)

	br11 := bufio.NewReader(fr11)
	line11, _, err := br11.ReadLine()
	check(err)

	fmt.Printf("\n11: %s", line11)

	//Write a program to remove newline characters from a file.
	//Read file, identify '\n' character, remove from slice, write to file.
	//TODO: It should be done moving the file cursor position.
	//TODO: Get file cursor for reading, get file cursor for writing, read until 
	//TODO: desired word, get position, place cursor before desired word, read and write from there,
	//TODO: close file cursors and truncate.
	//TODO: https://stackoverflow.com/a/2329972/3621080
	
	fr12, err := os.Open("./oof")
	check(err)
	defer fr12.Close()

	fw12, err := os.Open("./oof", os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer fw12.Close()






	//data12, err := ioutil.ReadFile("./oof")	//Returns []byte, err
	//check(err)
	//for i, e12 := range(data12) {
	//	if (e12 == '\n'){					//Identifying newline character(\x0a)
	//		remove(data12, i)				//Removing index i
	//	} 
	//}
	//f12, err := os.OpenFile("./oof-1", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//check(err)
	//defer f12.Close()
	//fmt.Fprintf(f12, string(data12))
	//fmt.Printf("\n12: %b", data12)

	//Write a program to remove a word from text file.
	//Read string, split string, find word, remove from slice(swaping current index to final index), dump to file.
	//TODO: It should be done moving the file cursor position.
	//TODO: https://stackoverflow.com/a/2329972/3621080
	br13, err := ioutil.ReadFile("./oof")
	check(err)
	nbr13 := strings.Replace(string(br13), "foo", "", -1)
	err = ioutil.WriteFile("./oof", []byte(nbr13), 0)
	check(err)
	
	//Write a program to rename a file using rename() function.
	err = os.Rename("./oof-1", "./foo-1")
	check(err)
	fmt.Println("\n14: Renaming file")

	//Write a program to count characters, words and lines in a text file.
	//Write a program to remove empty lines from a text file.
	//Write a program to remove specific line from a text file.
	//Write a program to replace specific line in a text file.
	//Write a program to read numbers from a file and write even, odd and prime numbers to separate file.
	//Write a program to convert uppercase to lowercase character and vice versa in a text file.
	//Write a program to compare two files.
	//Write a program to combine each line from first file with the corresponding line in second file.
	//Write a program to find and replace a word in a text file.
	//Write a program to print source code of same program.
	//Write a program to check if a file or directory exists.
	//Write a program to list all files and sub-directories recursively.
}
