package day3

import (
	"bufio"
	"fmt"
	"os"
)

// ParseTreesAndSquares => reads input file and returns a 2d array where trees are 1's and open squares are 0's
func ParseTreesAndSquares(path string) [][]int {
	filename := fmt.Sprintf("%s/day3/day3.txt", path)
	//filename := fmt.Sprintf("%s/day3/test.txt", path)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	var route [][]int
	for scanner.Scan() {
		row := []int{}
		line := scanner.Text()
		for _, r := range line {
			if r == rune('#') {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		route = append(route, row)

	}
	return route
}

// countTreesOnSlope => returns number of trees encountered on a given slope (represented by each move i.e. right and down)
func countTreesOnSlope(route [][]int, right int, down int) int {
	i := 0 // row i.e x-coordinate for top left
	step := 0
	trees := 0
	for i < len(route) {
		trees += route[i][(step*right)%len(route[i])] // since the row keeps on repeating endlessly
		i += down
		step++
	}
	fmt.Printf("Trees found on slope right %d and down %d is %d\n", right, down, trees)
	return trees
}

// ProductOfTreesOnSlopes => returns product of trees found on each slope
func ProductOfTreesOnSlopes(route [][]int, slopes [][]int) int {
	trees := 1
	for _, slope := range slopes {
		right, down := slope[0], slope[1]
		trees *= countTreesOnSlope(route, right, down)
	}
	return trees
}
