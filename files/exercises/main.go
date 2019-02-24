package main

import (
	"bufio"		//Package to read/write buffering
	"fmt"
	"io/ioutil" //Package to read files
	"os"
	//"io"
	"sort"
	"strings"
	"math/rand"
)

func check(e error) {
	if e != nil {
		os.Exit(1)
	}
}

//Here you separate men from boys
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
	fmt.Printf("\n1: ", string(data))

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
	for br6.Scan(){
		v6 = append(v6, br6.Text())
		check(br6.Err())
	}
	fmt.Printf("\n6: len %d, cap %d", len(v6), cap(v6))
	var tmp6 int
	for i, p := range(v6){
		fmt.Printf("\n6: line %d message: %s", i, p)	
		tmp6 = i
	}
	fmt.Printf("\n6: Total of lines: %d\n", tmp6+1) //i starts at 0
	
	//Write a program to find the longest words.
	//ReadAll, break in pieces, convert to string, sort array, get the biggest array.
	f7, err := os.Open("./foo")
	check(err)
	defer f7.Close()
	br7, err := ioutil.ReadAll(f7) //returns a byte slice
	check(err)
	var str7 string
	for _, p7 := range(br7){
		str7 = str7 + string(p7)
	}
	split7 := strings.Split(str7, " ") 
	sort.Strings(split7) //Sorts increasing order
	fmt.Printf("\n8: len %d, cap %d", len(split7), cap(split7))
	fmt.Printf("\n8: first: %s, len: %d\n%s", split7[0], len(split7[0]), str7)
	
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
	br9, err := ioutil.ReadFile("./foo")
	check(err)
	
	f9, err := os.OpenFile("./foobar", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f9.Close()
	fmt.Fprintf(f9, string(br9))

	//Write a program to count the frequency of words in a file.
	//Read file, create string slice, make a map of string: int, insert values in map.
	br10, err := ioutil.ReadFile("./foo")
	check(err)
	var str10 string
	for _, s10 := range(br10){
		str10 = str10 + string(s10)
	}
	vstr10 := strings.Fields(str10)

	map10 := make(map[string]int)
	for _, e10 := range(vstr10){
		_, ok := map10[e10]
		if ok == false {
			map10[e10] = 1
		} else {
			map10[e10]++  	
		}
	}
	for k10, v10 := range map10{
		fmt.Printf("\n10: key=%s, value=%d", k10, v10)
	}

	//Write a program to read a random line from a file.
	f11, err := os.Open("./foo")
	check(err)
	defer f11.Close()
	n11 := rand.Int63()
	m11 := rand.Int()
	f11.Seek(n11, m11)
	br11 := bufio.NewReader(f11)
	line11, _, err := br11.ReadLine()
	check(err)
	fmt.Printf("\n11: %s", line11)

	//Write a program to remove newline characters from a file.
	//Read file, identify '\n' character, remove from slice, write to file.
	data12, err := ioutil.ReadFile("./oof")	//Returns []byte, err
	check(err)
	for i, e12 := range(data12) {
		if (e12 == '\n'){					//Identifying newline character(\x0a)
			remove(data12, i)				//Removing index i
		}
	}
	f12, err := os.OpenFile("./oof-1", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f12.Close()
	fmt.Fprintf(f12, string(data12))
	fmt.Printf("\n12: %b", data12)

	//Write a program to read file contents and display on console.
	//Write a program to compare two files.
	//Write a program to combine each line from first file with the corresponding line in second file.
	//Write a program to read numbers from a file and write even, odd and prime numbers to separate file.
	//Write a program to count characters, words and lines in a text file.
	//Write a program to remove a word from text file.
	//Write a program to remove specific line from a text file.
	//Write a program to remove empty lines from a text file.
	//Write a program to count occurrences of a word in a text file.
	//Write a program to find and replace a word in a text file.
	//Write a program to replace specific line in a text file.
	//Write a program to print source code of same program.
	//Write a program to convert uppercase to lowercase character and vice versa in a text file.
	//Write a program to check if a file or directory exists.
	//Write a program to rename a file using rename() function.
	//Write a program to list all files and sub-directories recursively.
}
