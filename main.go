package main

import (
	"aoc/day1"
	"aoc/parseinput"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code")
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("problem fetching working dir")
		fmt.Println(err)
		os.Exit(1)
	}
	numbers := parseinput.FetchIntsFromFile(path, 1)
	fmt.Println("product of two sum is: ", day1.ProductOfSums(2, numbers))
	fmt.Println("product of two sum is: ", day1.ProductOfSums(3, numbers))

}
