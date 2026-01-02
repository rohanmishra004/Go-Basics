package main

import (
	"fmt"
	"net/http"
	"time"
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
	//if we use this once then the main routine will exit after getting response from one of the link but since we require it to execute for all the links we can use for loop to execute till it receives all the data

	// fmt.Println(<-c)

	//this for loop is an infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// 	// fmt.Println(<-c)
	// }

	//Alternatively in order to pause the program after every fetch call we can use sleep in time package. Sleep pauses the curent go routine, a -ve or 0 duration causes sleep to return immediately
	for l := range c {
		//if we put sleep in our main routine then we will cause all our messages to be throttled so in order to prevent that . In order to solve it we can use Function literals , function literals are similar to anonymous functions in JS or Lambda functions in Python
		// time.Sleep(2 * time.Second)
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

//We can make use of go routines to make these calls run in parallel

// we need to update the function as well here
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Might be down")
		// c <- "might be down"
		c <- link
	}
	fmt.Println(link, " is up!!!")
	// c <- "link is up"
	c <- link
}
