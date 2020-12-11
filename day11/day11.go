package day11

import (
	"bufio"
	"fmt"
	"os"
)

// SeatRule => function which takes in seatingMatrix and row, col of seat and returns numnber of occupied seats
type SeatRule func([][]rune, int, int) int

func ParseSeats(path string) [][]rune {
	filename := fmt.Sprintf("%s/day11/day11.txt", path)
	//filename := fmt.Sprintf("%s/day11/test.txt", path)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	var seatMatrix [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		seatMatrix = append(seatMatrix, []rune(scanner.Text()))
	}
	return seatMatrix
}

// CountOccupiedSeats => applies seating rule changes till no changes happen and then counts number of occupied seats
func CountOccupiedSeats(seatMatrix [][]rune, occupyLimit int, rule SeatRule) int {
	updated, changed := ApplySeatingRules(seatMatrix, occupyLimit, rule)
	for changed {
		updated, changed = ApplySeatingRules(updated, occupyLimit, rule)
	}
	occupied := 0
	for i := 0; i < len(updated); i++ {
		for j := 0; j < len(updated[0]); j++ {
			if updated[i][j] == rune('#') {
				occupied++
			}
		}
	}
	return occupied
}

// ApplySeatingRules => takes in a seating matrix and returns a new matrix after applying a round of seating rules
func ApplySeatingRules(seatMatrix [][]rune, occupyLimit int, rule SeatRule) ([][]rune, bool) {
	updated := [][]rune{}
	changed := false
	for i := 0; i < len(seatMatrix); i++ {
		row := []rune{}
		for j := 0; j < len(seatMatrix[0]); j++ {
			occupied := rule(seatMatrix, i, j)
			switch seatMatrix[i][j] {
			case rune('.'):
				row = append(row, seatMatrix[i][j])
			case rune('#'):
				if occupied >= occupyLimit {
					row = append(row, rune('L'))
					changed = true
				} else {
					row = append(row, seatMatrix[i][j])
				}
			case rune('L'):
				if occupied == 0 {
					row = append(row, rune('#'))
					changed = true
				} else {
					row = append(row, seatMatrix[i][j])
				}
			}
		}
		updated = append(updated, row)
	}
	return updated, changed
}

// CheckAdjacentSeats => returns number of occupied seats in adjacent area
func CheckAdjacentSeats(seatMatrix [][]rune, row int, col int) int {
	occupied := 0
	for i := max(row-1, 0); i >= 0 && i <= len(seatMatrix)-1 && i <= row+1; i++ {
		for j := max(col-1, 0); j >= 0 && j <= len(seatMatrix[0])-1 && j <= col+1; j++ {
			if i == row && j == col {
				continue
			}
			if seatMatrix[i][j] == rune('#') {
				occupied++
			}
		}
	}
	return occupied
}

// CheckFirstSeatInEachDirection => returns number of occupied seats when considering first visible seat in each direction
// person can see over floor but not over seats (empty or occupied)
func CheckFirstSeatInEachDirection(seatMatrix [][]rune, row int, col int) int {
	occupied := 0
	// We need to check 8 diff. directions
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			occupied += checkDirection(seatMatrix, row, col, i, j)
		}
	}
	return occupied
}

// checkDirection => keeps expanding the search in a given direction till we see an occupied or empty seat
func checkDirection(seatMatrix [][]rune, row int, col int, rowGrad int, colGrad int) int {
	i, j := row+rowGrad, col+colGrad
	for i >= 0 && i < len(seatMatrix) && j >= 0 && j < len(seatMatrix[0]) {
		if seatMatrix[i][j] == rune('#') {
			return 1

		}
		if seatMatrix[i][j] == rune('L') {
			return 0
		}
		i += rowGrad
		j += colGrad
	}
	return 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
