package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	ins   string
	value int
}

func ParseBootInstructions(path string) []Instruction {
	//filename := fmt.Sprintf("%s/day8/day8.txt", path)
	filename := fmt.Sprintf("%s/day8/test.txt", path)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	var instructions []Instruction
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		ins, valueStr := line[0], line[1]
		value, _ := strconv.Atoi(valueStr[1:])
		if valueStr[0] == '-' {
			value *= -1
		}
		instructions = append(instructions, Instruction{ins: ins, value: value})
	}
	return instructions
}

func FindAccumlatorBeforeRepeat(instructions []Instruction) int {
	seen := make(map[int]bool)
	i, acc := 0, 0
	for {
		if seen[i] {
			return acc
		}
		seen[i] = true
		switch instructions[i].ins {
		case "acc":
			acc += instructions[i].value
			i++
		case "jmp":
			i += instructions[i].value
		case "nop":
			i++
		}
	}
}
