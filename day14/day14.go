package day14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BitInstruction struct {
	onesMask     int
	zerosMask    int
	addressMasks []int
	memory       map[int]int
}

type InstructionProcessor func([]BitInstruction) map[int]int

func ParseInstructions(path string) []string {
	//filename := fmt.Sprintf("%s/day14/day14.txt", path)
	//filename := fmt.Sprintf("%s/day14/test.txt", path)
	filename := fmt.Sprintf("%s/day14/test2.txt", path)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	return instructions
}

func ParseBitInstructionsV1(instructions []string) []BitInstruction {
	var bitInstructions []BitInstruction
	var bitInstruction BitInstruction
	masks := regexp.MustCompile(`mask = ([\d\w]+)`)
	mems := regexp.MustCompile(`mem\[([\d]+)\] = ([\d]+)`)
	for _, line := range instructions {
		if strings.HasPrefix(line, "mask") {
			maskStrs := masks.FindStringSubmatch(line)
			onesMask, zerosMask := createBitmasks(maskStrs[1])
			bitInstruction = BitInstruction{onesMask: onesMask, zerosMask: zerosMask, memory: make(map[int]int)}
			bitInstructions = append(bitInstructions, bitInstruction)
		} else {
			memInstruction := mems.FindStringSubmatch(line)
			updateMemoryMap(memInstruction[1], memInstruction[2], bitInstruction.memory)
		}
	}
	return bitInstructions
}

func ParseBitInstructionsV2(instructions []string) []BitInstruction {
	var bitInstructions []BitInstruction
	var bitInstruction BitInstruction
	masks := regexp.MustCompile(`mask = ([\d\w]+)`)
	mems := regexp.MustCompile(`mem\[([\d]+)\] = ([\d]+)`)
	for _, line := range instructions {
		if strings.HasPrefix(line, "mask") {
			if len(bitInstruction.memory) > 0 {
				updateMemoryMapWithAddressMasks(bitInstruction.memory, bitInstruction.addressMasks)
			}
			maskStrs := masks.FindStringSubmatch(line)
			addressMasks := createAddressMasks(maskStrs[1])
			bitInstruction = BitInstruction{addressMasks: addressMasks, memory: make(map[int]int)}
			bitInstructions = append(bitInstructions, bitInstruction)
		} else {
			memInstruction := mems.FindStringSubmatch(line)
			updateMemoryMap(memInstruction[1], memInstruction[2], bitInstruction.memory)
		}
	}
	fmt.Println("Bit instructions V2: ", bitInstructions)
	return bitInstructions
}

func updateMemoryMap(keyStr string, valueStr string, memory map[int]int) {
	key, _ := strconv.Atoi(keyStr)
	value, _ := strconv.Atoi(valueStr)
	memory[key] = value
}

func updateMemoryMapWithAddressMasks(memory map[int]int, addressMasks []int) {
	var newKey int
	originalKeys := []int{}
	for k := range memory {
		originalKeys = append(originalKeys, k)
	}
	for k, v := range memory {
		for _, mask := range addressMasks {
			newKey = k | mask
			memory[newKey] = v
		}
	}
	for _, k := range originalKeys {
		delete(memory, k)
	}
	fmt.Println("Updated memory", memory, addressMasks)
}

func createBitmasks(bitStr string) (int, int) {
	var onesMask, zerosMask int = 0, 0
	j := len(bitStr) - 1
	for i, v := range bitStr {
		if v == rune('0') {
			continue
		} else if v == rune('1') {
			onesMask |= 1 << (j - i)
		}
		zerosMask += 1 << (j - i)
	}
	return onesMask, zerosMask
}

func createAddressMasks(bitStr string) []int {
	mask, j := 0, len(bitStr)-1
	addressMasks := []int{0}
	masks := []int{}
	for i, v := range bitStr {
		masks = []int{}
		for k := 0; k < len(addressMasks); k++ {
			mask = addressMasks[k]
			mask |= 1 << (j - i)
			if v == rune('X') {
				masks = append(masks, mask)
				mask ^= 1 << (j - i)
				masks = append(masks, mask)
			} else if v == rune('1') {
				masks = append(masks, mask)
			} else {
				mask ^= 1 << (j - i)
				masks = append(masks, mask)
			}
		}
		addressMasks = make([]int, len(masks))
		copy(addressMasks, masks)
	}
	return addressMasks

}

func ApplyInstructionsV1(instructions []BitInstruction) map[int]int {
	results := make(map[int]int)
	for _, instruction := range instructions {
		for key, value := range instruction.memory {
			value |= instruction.onesMask
			value &= instruction.zerosMask
			results[key] = value
		}
	}
	return results
}

func ApplyInstructionsV2(instructions []BitInstruction) map[int]int {
	results := make(map[int]int)
	for _, instruction := range instructions {
		for key, value := range instruction.memory {
			results[key] = value
		}
	}
	return results
}

func SumOfMemory(instructions []BitInstruction, processor InstructionProcessor) int {
	updatedMemory := processor(instructions)
	sum := 0
	for _, v := range updatedMemory {
		sum += v
	}
	return sum
}
