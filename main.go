package main

import (
	"aoc/day13"
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
	//baggageRules := day7.ParseBaggageRules(path)
	//reverseGraph := day7.ReverseBaggageGraph(baggageRules)
	//fmt.Println("Number of valid outer bags is: ", day7.CountOuterBags("shiny gold", reverseGraph))
	//fmt.Println("Number of inner bags: ", day7.CountBagsInside("shiny gold", baggageRules))

	// Day 8
	//instructions := day8.ParseBootInstructions(path)
	//fmt.Println("Accumulator value before loop is: ", day8.FindAccumlatorBeforeRepeat(instructions))

	// Day 9
	//numbers := day9.ParseNumbers(path)
	//invalid := day9.FindFirstInValidNumber(numbers, 25)
	//index, invalidNumber := invalid[0], invalid[1]
	//fmt.Println(index, invalidNumber)
	//fmt.Println("XMAS weakness is: ", day9.FindEncryptionWeakness(numbers[:index], invalidNumber))

	// Day 10
	//joltages := day10.ParseJoltages(path)
	//fmt.Println("Distribution of input joltages is: ", day10.FindJoltageDistribution(joltages))
	//fmt.Println("Number of distinct arrangements of joltages: ", day10.FindDistinctJoltagePaths(joltages))

	// Day 11
	//seatMatrix := day11.ParseSeats(path)
	//fmt.Println("Part 1: Number of occupied seats is: ", day11.CountOccupiedSeats(seatMatrix, 4, day11.CheckAdjacentSeats))
	//fmt.Println("Part 2: Number of occupied seats is: ", day11.CountOccupiedSeats(seatMatrix, 5, day11.CheckFirstSeatInEachDirection))

	// Day 12
	//navigations := day12.ParseDirections(path)
	//initialPosition := day12.Point{}
	// finalPosition := day12.FollowDirections(navigations)
	// fmt.Println("Part1: Manhattan distance between start and finish is: ", day12.ManhattanDistance(initialPosition, finalPosition))
	// newFinalPosition := day12.FollowDirectionsRelativeToWayPoint(navigations, day12.Point{X: 10, Y: -1})
	// fmt.Println("Part1: Manhattan distance between start and finish is: ", day12.ManhattanDistance(initialPosition, newFinalPosition))

	// Day 13
	start, busTimes := day13.ParseBusTimes(path)
	//fmt.Println(start, busTimes)
	fmt.Println("Product of earliest departing busID and waiting time is: ", day13.ProductOfEarliestDepartingBusAndWaitingTime(start, busTimes))
	fmt.Println("Ideal starting time so that all buses can leave with time interval equal to the offset in bus times array: ", day13.FindIdealStart(busTimes))
	//fmt.Println(day13.FindIdealStart([]int{13, 0, 0, 59, 0, 31, 19}))

	//fmt.Println(day13.Egcd(11, 8))
}
