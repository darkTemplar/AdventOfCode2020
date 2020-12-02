package parseinput

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var baseURL = "https://adventofcode.com/2019/day/%d/input"

// FetchDailyInputFromWeb... fetches daily input from webpage (currently not available from AOC)
func FetchDailyInputFromWeb(day int) {
	fmt.Println("Parsing input for day: ", day)
	inputURL := fmt.Sprintf(baseURL, day)
	fmt.Println("Input url is: ", inputURL)
	resp, err := http.Get(inputURL)
	if err != nil {
		fmt.Println("Something went wrong fetching input for day: ", day)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request to fetch input failed")
		return
	}
	fmt.Println(resp.Body)
}

// FetchIntsFromFile ... reads line by line input from file of naming format day<num>.txt and returns an array of integers
func FetchIntsFromFile(currentDir string, day int) []int {
	filename := fmt.Sprintf("%s/day%d/day%d.txt", currentDir, day, day)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening input file")
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	var results []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		results = append(results, number)
	}
	return results
}
