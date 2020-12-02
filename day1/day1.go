package day1

import (
	"sort"
)

// TwoSum => finds 2 numbers in the sorted array which sum to 2020
func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for right > left {
		if numbers[left]+numbers[right] == target {
			return []int{numbers[left], numbers[right]}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

// ThreeSum => finds 3 numbers in the sorted array which sum to 2020
func ThreeSum(numbers []int, target int) []int {
	for i, number := range numbers {
		pair := TwoSum(numbers[i+1:], target-number)
		if len(pair) > 0 {
			pair = append(pair, number)
			return pair
		}
	}
	return []int{}
}

// ProductOfSums => returns product of sums. sums depends on the level argument passed in.
// E.g. if its 2, then sums will contain 2 numbers
func ProductOfSums(level int, numbers []int) int {
	product := 1
	sort.Ints(numbers)
	var results []int
	switch level {
	case 2:
		results = TwoSum(numbers, 2020)
	case 3:
		results = ThreeSum(numbers, 2020)
	default:
		return 1
	}
	for _, number := range results {
		product *= number
	}
	return product
}
