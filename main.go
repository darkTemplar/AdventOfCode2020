package main

import (
	"aoc/day7"
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
	//policies := day2.FetchPolicyAndPassword(path, 2)
	//fmt.Println("Number of valid passwords is as per first policy is: ", day2.CountValidPasswords(policies, 1))
	//fmt.Println("Number of valid passwords is as per first policy is: ", day2.CountValidPasswords(policies, 2))

	//Day 3
	//route := day3.ParseTreesAndSquares(path)
	//slope := []int{3, 1}
	//fmt.Println("Number of trees on slopes: ", day3.ProductOfTreesOnSlopes(route, [][]int{slope}))
	//slopes := [][]int{[]int{1, 1}, []int{3, 1}, []int{5, 1}, []int{7, 1}, []int{1, 2}}
	//fmt.Println("Number of trees on slopes: ", day3.ProductOfTreesOnSlopes(route, slopes))

	//Day 4
	//passports := day4.ParsePassportData(path)
	//fmt.Println("Number of valid passports is: ", day4.CountValidPassports(passports))

	// Day 5
	//passes := day5.ParseBoardingPasses(path)
	//fmt.Println("Maximum seatID is: ", day5.MaxSeatID(passes))
	//fmt.Println("Missing seatID is: ", day5.FindMissingSeat(passes))

	// Day 6
	//answerMaps := day6.ParseCustomFormInput(path)
	//fmt.Println("total answer count is: ", day6.CountAnswersFromGroups(answerMaps))
	//fmt.Println("total all yes question count is: ", day6.CountAllYesQuestions(answerMaps))

	// Day 7
	baggageRules := day7.ParseBaggageRules(path)
	reverseGraph := day7.ReverseBaggageGraph(baggageRules)
	fmt.Println("Number of valid outer bags is: ", day7.CountOuterBags("shiny gold", reverseGraph))
	fmt.Println("Number of inner bags: ", day7.CountBagsInside("shiny gold", baggageRules))

}
