package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	//here we create an empty byte of size 999999(assumption)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//we can replace the whole code above with these simple lines -

	//io.Copy take two values(writeInterface , readerInterface) so here - os.Stdout - writeInterface , resp.Body- readInterface

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("JUST WROTE THIS MANY BYTES ::", len(bs))
	return len(bs), nil
}
