package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PasswordPolicy => Password and policy storage
type PasswordPolicy struct {
	Min      int
	Max      int
	Required rune
	Password string
}

func isCharInRange(pwd string, target rune, min int, max int) int {
	charFreq := make(map[rune]int)
	for _, char := range pwd {
		if _, ok := charFreq[char]; ok {
			charFreq[char]++
		} else {
			charFreq[char] = 1
		}
	}
	freq, ok := charFreq[target]
	if !ok || (freq < min || freq > max) {
		return 0
	}
	return 1
}

func isCharAtRequiredIndexes(pwd string, target rune, min int, max int) int {
	runes := []rune(pwd)
	totalOccurences := 0
	if len(runes) >= min && runes[min-1] == target {
		totalOccurences++
	}
	if len(runes) >= max && runes[max-1] == target {
		totalOccurences++
	}

	if totalOccurences == 1 {
		return 1
	}
	fmt.Println("Invalid")
	return 0
}

// CountValidPasswords => count the number of valid passwords as per specified policy type
func CountValidPasswords(policies []PasswordPolicy, policyType int) int {
	valid := 0
	var validationFunc func(string, rune, int, int) int
	if policyType == 1 {
		validationFunc = isCharInRange
	} else {
		validationFunc = isCharAtRequiredIndexes
	}
	for _, policy := range policies {
		valid += validationFunc(policy.Password, policy.Required, policy.Min, policy.Max)
	}
	return valid
}

// FetchPolicyAndPassword => extract policy and passwords from input file
func FetchPolicyAndPassword(currentDir string, day int) []PasswordPolicy {
	filename := fmt.Sprintf("%s/day%d/day%d.txt", currentDir, day, day)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening input file")
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	var results []PasswordPolicy
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		requiredChar := []rune(items[1])
		policy := PasswordPolicy{Password: items[2], Required: requiredChar[0]}
		minmax := strings.Split(items[0], "-")
		policy.Min, _ = strconv.Atoi(minmax[0])
		policy.Max, _ = strconv.Atoi(minmax[1])
		results = append(results, policy)
	}
	return results
}
