package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ParseNumbers(path string) []int {
	filename := fmt.Sprintf("%s/day9/day9.txt", path)
	//filename := fmt.Sprintf("%s/day9/test.txt", path)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	var numbers []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	return numbers
}

func FindFirstInValidNumber(numbers []int, preamble int) []int {

	for i, j := 0, preamble; j < len(numbers); i, j = i+1, j+1 {
		if !hasTwoSum(numbers[i:j], numbers[j]) {
			return []int{j, numbers[j]}
		}
	}
	return []int{-1, -1}
}

func hasTwoSum(nums []int, target int) bool {
	pairs := make(map[int]bool)
	for _, num := range nums {
		if pairs[target-num] {
			return true
		}
		pairs[num] = true
	}
	return false
}

func FindEncryptionWeakness(numbers []int, target int) int {
	sum := 0
	contiguous := []int{}
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
		contiguous = append(contiguous, numbers[i])
		if sum == target {
			fmt.Println("target found")
			return SumMinAndMax(contiguous)
		}
		if sum > target {
			for j := 0; j < i; j++ {
				sum -= contiguous[0]
				contiguous = contiguous[1:]
				if sum == target {
					return SumMinAndMax(contiguous)
				}
				if sum < target {
					break
				}
			}
		}
	}
	return -1
}

func SumMinAndMax(numbers []int) int {
	min, max := math.MaxInt32, math.MinInt32
	for _, number := range numbers {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}
	return min + max
}
