/**
Create a program that reads the contents of text file and prints it on the terminal

The file to open should be provided as an argument . ex if we run using go run main.go myfile.txt it should open myfile.txt

To read in the arguments provided to a program you can reference the variable os.Args which is a slice of string

To open a file, check out the documentation

What interface does the file type implement

If the file type implements Reader interface we might need to use io.Copy function
**/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args) //this will provide us the list of all the command line arguments

	fmt.Println(os.Args[1]) //this will give us the file name which we entered in the argument

	//to read the file value we can use as
	file := os.Args[1] //this gives us the name of the file

	f, err := os.Open(file) //this will return a pointer to a file and an error if one occured
	if err != nil {
		fmt.Println("Error::", err)
		os.Exit(1)
	}

	//the file function directly implements a read func so we can use io.Copy to read the data and write it to the terminal
	io.Copy(os.Stdout, f)

}
