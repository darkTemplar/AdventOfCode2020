package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// ParseBoardingPasses => parse input file for boarding pass data and return array of seatIDs
func ParseBoardingPasses(path string) []int {
	filename := fmt.Sprintf("%s/day5/day5.txt", path)
	//filename := fmt.Sprintf("%s/day5/test.txt", path)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	var passes []int
	for scanner.Scan() {
		row, column := 0, 0
		content := scanner.Text()
		for i, r := range content {
			switch r {
			case rune('F'):
				row += 0
			case rune('B'):
				row += int(math.Exp2(float64(6 - i)))
			case rune('R'):
				column += int(math.Exp2(float64(9 - i)))
			case rune('L'):
				column += 0
			}
		}
		passes = append(passes, row*8+column)
	}
	return passes
}

// MaxSeatID => finds maximum seatID in array
func MaxSeatID(passes []int) int {
	maxID := 0
	for _, seatID := range passes {
		if seatID > maxID {
			maxID = seatID
		}
	}
	return maxID
}

// FindMissingSeat => find the only missing seatID in the passes array (the input is continuous except for missing seatID)
func FindMissingSeat(passes []int) int {
	sort.Ints(passes)
	min, max := passes[0], passes[len(passes)-1]
	// calculate
	apSum := ((min + max) * (len(passes) + 1)) / 2
	for _, seatID := range passes {
		apSum -= seatID
	}
	return apSum
}
