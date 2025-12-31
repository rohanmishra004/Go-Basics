package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf75e",
		"white": "#ffffff",
	}

	// fmt.Println(colors)
	printMap(colors)
}

//we want to write a new function that accepts a map and iterate over the map and print out all the key values pairs

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
