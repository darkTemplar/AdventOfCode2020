package day7

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BaggageRules map[string]map[string]int

func ParseBaggageRules(path string) BaggageRules {
	filename := fmt.Sprintf("%s/day7/day7.txt", path)
	//filename := fmt.Sprintf("%s/day7/test.txt", path)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`bags|bag`)
	baggageRules := make(map[string]map[string]int)
	var key string
	for scanner.Scan() {
		line := scanner.Text()
		rules := re.Split(line, -1)                   // return all matches
		for idx, rule := range rules[:len(rules)-1] { // we exclude last value as its blank
			rule = strings.TrimSpace(rule)
			components := strings.Split(rule, " ")
			strlen := len(components)
			// zero index means the top level container bag
			// other indices represent the bags contained within the top level bag
			if idx == 0 {
				// color indexes will be the last 2 elems of slice
				key = strings.Join(components[strlen-2:], " ")
				//fmt.Println("Outer Baggage key is: ", key)
				updateBaggageRule(key, baggageRules)
			} else {
				innerKey := strings.Join(components[strlen-2:], " ")
				// ignore if no inner bags specified in the rule
				if innerKey == "no other" {
					continue
				}
				value, _ := strconv.Atoi(components[strlen-3])
				//fmt.Println("Inner Baggage key & value is: ", innerKey, value)
				updateBaggageRuleKey(innerKey, value, baggageRules[key])
			}
		}
	}
	return baggageRules
}

func updateBaggageRule(key string, rules BaggageRules) {
	if _, ok := rules[key]; !ok {
		rules[key] = make(map[string]int)
	}
}

func updateBaggageRuleKey(key string, value int, rule map[string]int) {
	if _, ok := rule[key]; ok {
		rule[key] += value
	} else {
		rule[key] = value
	}
}

func ReverseBaggageGraph(baggageRules BaggageRules) map[string]map[string]bool {
	reverseGraph := make(map[string]map[string]bool)
	for outerbag := range baggageRules {
		reverse(outerbag, baggageRules, reverseGraph)
	}
	return reverseGraph
}

func reverse(src string, graph BaggageRules, reverseGraph map[string]map[string]bool) {
	innerBags, _ := graph[src]
	for innerBag := range innerBags {
		if _, ok := reverseGraph[innerBag]; ok {
			if !reverseGraph[innerBag][src] {
				reverseGraph[innerBag][src] = true
			}
		} else {
			reverseGraph[innerBag] = map[string]bool{src: true}
		}
		reverse(innerBag, graph, reverseGraph)
	}
}

func CountOuterBags(targetBag string, reverseGraph map[string]map[string]bool) int {
	// calculate how many nodes are reachable from "shiny gold" node now
	count := 0
	frontier, seen := []string{targetBag}, map[string]bool{targetBag: true}
	for len(frontier) > 0 {
		node := frontier[0]
		frontier = frontier[1:]
		if neighbors, ok := reverseGraph[node]; ok {
			for neighbor := range neighbors {
				if seen[neighbor] {
					continue
				}
				seen[neighbor] = true
				frontier = append(frontier, neighbor)
				count++
			}
		}
	}
	return count
}

func CountBagsInside(targetBag string, baggageRules BaggageRules) int {
	// calculate how many nodes are reachable from "shiny gold" node now
	count := 0
	if innerBags, _ := baggageRules[targetBag]; len(innerBags) > 0 {
		for innerBag, number := range innerBags {
			bagsInside := CountBagsInside(innerBag, baggageRules)
			// 1 bag inside indicates terminal level i.e. the bag has no other inside bags
			if bagsInside == 1 {
				count += number * bagsInside
			} else { // count the bag containing inside bags
				count += (number * bagsInside) + number
			}
		}
		return count
	}
	return 1
}
