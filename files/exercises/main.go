package main

import (
	"bufio"
	"fmt"
	"io/ioutil" //Package to read files
	"os"
	//"io"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		os.Exit(1)
	}
}

//Here you separate men from boys
//ETA 30h

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
	var n4 int64 = 10
	f4, err := os.Open("./foo")
	check(err)
	defer f4.Close()
	f4a, err := os.Stat("./foo")
	check(err)
	f4.Seek(f4a.Size()-n4, 0)	//Reads from total bytes - n
	//buffer4 := make([]byte, n4)									//Long way to do the same
	//br4, err := io.ReadFull(f4, 									//Long way to do the samebuffer4)
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
	
	//Wirte a program to write a list to a file.
	//Write a program to copy the contents of a file to another file.
	//Write a program to count the frequency of words in a file.
	//Write a program to combine each line from first file with the corresponding line in second file.
	//Write aprogram to read a random line from a file.
	//Write a program to assess if a file is closed or not.
	//Write a program to remove newline characters from a file.
	//Write a program to create a file and write contents, save and close the file.
	//Write a program to read file contents and display on console.
	//Write a program to read numbers from a file and write even, odd and prime numbers to separate file.
	//Write a program to compare two files.
	//Write a program to copy contents from one file to another file.
	//Write a program to merge two file to third file.
	//Write a program to count characters, words and lines in a text file.
	//Write a program to remove a word from text file.
	//Write a program to remove specific line from a text file.
	//Write a program to remove empty lines from a text file.
	//Write a program to find occurrence of a word in a text file.
	//Write a program to count occurrences of a word in a text file.
	//Write a program to count occurrences of all words in a text file.
	//Write a program to find and replace a word in a text file.
	//Write a program to replace specific line in a text file.
	//Write a program to print source code of same program.
	//Write a program to convert uppercase to lowercase character and vice versa in a text file.
	//Write a program to check if a file or directory exists.
	//Write a program to rename a file using rename() function.
	//Write a program to list all files and sub-directories recursively.
}
