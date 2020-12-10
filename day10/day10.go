package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ParseJoltages(path string) []int {
	filename := fmt.Sprintf("%s/day10/day10.txt", path)
	//ilename := fmt.Sprintf("%s/day10/test.txt", path)
	//filename := fmt.Sprintf("%s/day10/test2.txt", path)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	joltages := []int{0} // since charging outlet has effective rating of 0
	for scanner.Scan() {
		joltage, _ := strconv.Atoi(scanner.Text())
		joltages = append(joltages, joltage)
	}
	return joltages
}

func FindJoltageDistribution(joltages []int) [3]int {
	sort.Ints(joltages)
	distribution := [3]int{0, 0, 0}
	inputRating := joltages[0]
	for _, joltage := range joltages[1:] {
		if joltage-inputRating > 3 {
			fmt.Println("No more compatible adapters left")
			return distribution
		}
		distribution[joltage-inputRating-1]++
		inputRating = joltage
	}
	// since your device is like the final adapter with rating +3 of highest rating, we increment distribution of diff 3 by 1
	distribution[2]++
	return distribution
}

func FindDistinctJoltagePaths(joltages []int) int {
	sort.Ints(joltages)
	paths := []int{1}
	var idx int
	for i, joltage := range joltages[1:] {
		idx = i + 1
		// init paths
		paths = append(paths, 0)
		for j := 1; idx-j >= 0 && j <= 3; j++ {
			if joltage-joltages[idx-j] > 3 {
				continue
			}
			paths[idx] += paths[idx-j]
		}
	}
	return paths[len(joltages)-1]
}
