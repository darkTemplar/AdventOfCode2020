package main

import (
	"aoc/day2"
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
	// Day 1: Solution
	//numbers := parseinput.FetchIntsFromFile(path, 1)
	//fmt.Println("product of two sum is: ", day1.ProductOfSums(2, numbers))
	//fmt.Println("product of two sum is: ", day1.ProductOfSums(3, numbers))

	// Day 2
	policies := day2.FetchPolicyAndPassword(path, 2)
	fmt.Println("Number of valid passwords is as per first policy is: ", day2.CountValidPasswords(policies, 1))
	fmt.Println("Number of valid passwords is as per first policy is: ", day2.CountValidPasswords(policies, 2))

}
