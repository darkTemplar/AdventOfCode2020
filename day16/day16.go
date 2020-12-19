package day16

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func ParseTickets(path string) ([]string, [][2]int, []int, [][]int) {
	filename := fmt.Sprintf("%s/day16/day16.txt", path)
	//filename := fmt.Sprintf("%s/day16/test.txt", path)
	//filename := fmt.Sprintf("%s/day16/test2.txt", path)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	contentsStr := string(contents)
	var intervals [][2]int
	intervalRe := regexp.MustCompile(`\d{1,3}-\d{1,3}`)
	fieldNamesRe := regexp.MustCompile(`([a-z ]+):`)
	// parse fieldNames
	matches := fieldNamesRe.FindAllStringSubmatch(contentsStr, -1)
	fieldNames := []string{}
	for _, match := range matches {
		fieldNames = append(fieldNames, match[1])
	}
	// remove ticket fields
	fieldNames = fieldNames[:len(fieldNames)-2]
	// parse intervals
	intervalArr := intervalRe.FindAllString(contentsStr, -1)
	for _, elem := range intervalArr {
		interval := [2]int{}
		intervalStr := strings.Split(elem, "-")
		interval[0], _ = strconv.Atoi(intervalStr[0])
		interval[1], _ = strconv.Atoi(intervalStr[1])
		intervals = append(intervals, interval)
	}

	// parse tickets
	ticketRe := regexp.MustCompile(`(your ticket):(\s+)([\d,]+)`)
	ticketsArr := ticketRe.FindStringSubmatch(contentsStr)
	yourTickets := []int{}
	for _, t := range strings.Split(ticketsArr[3], ",") {
		ticket, _ := strconv.Atoi(t)
		yourTickets = append(yourTickets, ticket)
	}
	nearbyRe := regexp.MustCompile(`(nearby tickets):(\s+)([\d\s,]+)`)
	nearbyArr := nearbyRe.FindStringSubmatch(contentsStr)
	nearby := strings.Split(nearbyArr[3], "\n")
	nearbyTickets := parseTicketsArray(nearby)
	return fieldNames, intervals, yourTickets, nearbyTickets
}

func parseTicketsArray(arr []string) [][]int {
	var tickets [][]int
	for _, line := range arr {
		ticketStr := strings.Split(strings.TrimSpace(line), ",")
		if len(ticketStr) < 2 {
			continue
		}
		arr := []int{}
		for _, t := range ticketStr {
			ticket, _ := strconv.Atoi(t)
			arr = append(arr, ticket)
		}
		tickets = append(tickets, arr)
	}
	return tickets
}

func CombineFieldsAndIntervals(fields []string, intervals [][2]int) map[string][][2]int {
	ranges := make(map[string][][2]int)
	for i := 0; i <= len(fields)-1; i++ {
		ranges[fields[i]] = [][2]int{intervals[2*i], intervals[2*i+1]}
	}
	return ranges
}

func findFieldOrdering(fieldMap map[string][][2]int, tickets [][]int) map[string]int {
	graph := make(map[string][]int)
	for i := 0; i < len(tickets[0]); i++ {
		// construct list of column values
		column := []int{}
		for j := 0; j < len(tickets); j++ {
			column = append(column, tickets[j][i])
		}
		// check if row satisfies anyone row
		for k, intervals := range fieldMap {
			if withinRange(intervals, column) {
				if edges, ok := graph[k]; ok {
					edges = append(edges, i)
				} else {
					graph[k] = []int{i}
				}
			}
		}
	}
	ordering := biPartiteMatching(graph)
	fmt.Println("Field Ordering: ", ordering)
	return ordering
}

//TODO: implement
func biPartiteMatching(graph map[string][]int) map[string]int {
	ordering := make(map[string]int)
	return ordering
}

func FindDepartureFieldsProduct(fieldMap map[string][][2]int, your []int, validNearby [][]int) int {
	product := 1
	ordering := findFieldOrdering(fieldMap, validNearby)
	for k, v := range ordering {
		fmt.Println("key name: ", k, v)
		if strings.HasPrefix(k, "departure") {
			product *= your[v]
		}
	}
	return product
}

func FindTicketScanningErrorRate(intervals [][2]int, tickets [][]int) (int, [][]int) {
	var errorRate int
	validTickets := [][]int{}
	merged := mergeIntervals(intervals)
	for _, row := range tickets {
		valid := true
		for _, t := range row {
			if !isValid(merged, t) {
				errorRate += t
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, row)
		}
	}
	return errorRate, validTickets
}

// mergeIntervals => takes a bunch of overlapping intervals and sorts and merges them
func mergeIntervals(intervals [][2]int) [][2]int {
	// sort intervals
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	start, end := intervals[0][0], intervals[0][1]
	var merged [][2]int
	for i := 1; i < len(intervals); i++ {
		// start of new interval is greater than end, means we start new interval
		if intervals[i][0] > end {
			merged = append(merged, [2]int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else { // expand current interval
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		}
	}
	// add last interval
	merged = append(merged, [2]int{start, end})
	return merged
}

// isValid => given a key, searches the intervals to check if key is present (range query)
func isValid(intervals [][2]int, key int) bool {
	low, high := 0, len(intervals)-1
	mid := (low + high) / 2
	for low >= 0 && low <= high && high < len(intervals) {
		if key >= intervals[mid][0] && key <= intervals[mid][1] {
			return true
		}
		if key > intervals[mid][1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = (low + high) / 2
	}
	return false
}

// wiythinRange => checks if all given values are within the intervals provided
func withinRange(intervals [][2]int, values []int) bool {
	for _, v := range values {
		if !isValid(intervals, v) {
			return false
		}
	}
	return true
}
