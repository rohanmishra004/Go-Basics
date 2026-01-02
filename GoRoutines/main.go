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
	//Inside the main func we will create our first channel to interact with different go routines

	//syntax for creating channel
	c := make(chan string)

	//create a new thread for go routine
	for _, link := range links {
		//since we have created the channel we need to pass it to the function
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

//We can make use of go routines to make these calls run in parallel

// we need to update the function as well here
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Might be down")
		c <- "might be down"
	}
	fmt.Println(link, " is up!!!")
	c <- "link is up"
}
