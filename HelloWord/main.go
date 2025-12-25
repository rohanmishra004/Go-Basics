/**
Basic program

Commands -

1 - go run file.go - this command is use to compile  run the go file

2 - go build file.go - this commands simply compiles the go code and creates an executable file with a .exe extension. it does not run the code. To run this exe file we can simply use ./filename and it runs the file

3 - go fmt - automatically format all the code in each in the current directory

4 - go install - compiles and install a package similar to npm install

4 - go get - download the raw source code of someone else's package

5 - go test - runs any tests associated with the current project


*/

//package is a collection of commonly used files. A package can have many related files
//Inside go there are two different types of packages - Executable type and Reusable type
//Executable type package are generally use to run something
//Reusuable packages - code used as 'helpers'. Good place to put reusable logic or helper function

//depending upon the name of the package which decides whether we are creating a executable file or a reusuable file. Ex - package.main --> go build --> main.exe

/**If we have used any other package name for our package it would have created a helper.
package anyything ---> go build --> nothing! - compiling a non-main package gives nothing

Basically if we are using main , then we are creating an executable file. If we are using any other name other than main in package we are creating a reusuable or a helper file
**/

/**
Also one important thing , if we are trying to make any file executable then we need to have main function there

**/

package main

import "fmt"

func main() {
	fmt.Println("Hi There!")
}
