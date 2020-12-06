package day6

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ParseCustomFormInput => parse out answers from custom form responses for each group
func ParseCustomFormInput(path string) []map[rune]int {
	filename := fmt.Sprintf("%s/day6/day6.txt", path)
	//filename := fmt.Sprintf("%s/day6/test.txt", path)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	var answerMaps []map[rune]int
	answerContents := strings.Split(string(contents), "\n")
	yes, respondents := make(map[rune]int), 0
	for _, answerGroup := range answerContents {

		// encountering whitespace means new group of answers has started
		runes := []rune(answerGroup)
		if len(runes) == 0 {
			yes[rune('#')] = respondents
			answerMaps = append(answerMaps, yes)
			// reset answer map and number of respondents
			yes = make(map[rune]int)
			respondents = 0
			continue
		}
		//keep track of number of respondents for each group
		respondents++
		for _, answer := range answerGroup {
			if _, ok := yes[answer]; ok {
				yes[answer]++
			} else {
				yes[answer] = 1
			}
		}
	}
	return answerMaps
}

// CountAnswersFromGroups => returns sum of all responses across all groups
func CountAnswersFromGroups(answerMaps []map[rune]int) int {
	count := 0
	for _, answerMap := range answerMaps {
		count += len(answerMap) - 1 // exclude the respondents key
	}
	return count
}

// CountAllYesQuestions => returns count of questions which got yes responses from every respondent in each group
func CountAllYesQuestions(answerMaps []map[rune]int) int {
	allYesCount := 0
	respondentKey := rune('#')
	for _, answerMap := range answerMaps {
		respondents, _ := answerMap[respondentKey]
		for _, count := range answerMap {
			if count == respondents {
				allYesCount++
			}
		}
	}
	// subtract length of answerMaps as in above loop we are always counting the respondentsKey as well
	return allYesCount - len(answerMaps)
}
