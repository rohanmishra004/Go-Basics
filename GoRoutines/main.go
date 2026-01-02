package main

import (
	"fmt"
	"net/http"
)

// we are checking the status of each of these websites

// in this code we are running a sequential check and therefore the code is a bit slow in execution
func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
	}
	//create a new thread for go routine
	for _, link := range links {
		go checkLink(link)
	}
}

//We can make use of go routines to make these calls run in parallel

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Might be down")
	}
	fmt.Println(link, " is up!!!")
}
