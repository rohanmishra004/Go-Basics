package main

import "fmt"

func main() {
	nums := []int{} //we create an empty array and then we use for loop to fill it with 10 values

	for i := 0; i < 11; i++ {
		nums = append(nums, i)
	}
	fmt.Println("NUMS::", nums)

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println("Even number", num)
		} else {
			fmt.Println("Odd Number", num)
		}
	}
}
