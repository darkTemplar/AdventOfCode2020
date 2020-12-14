package day13

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func ParseBusTimes(path string) (int, []int) {
	var startTime int
	var busTimes []int

	filename := fmt.Sprintf("%s/day13/day13.txt", path)
	//filename := fmt.Sprintf("%s/day13/test.txt", path)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}

	schedule := strings.Split(string(contents), "\n")
	startTime, _ = strconv.Atoi(schedule[0])
	times := strings.Split(schedule[1], ",")
	for _, t := range times {
		if t == "x" {
			busTimes = append(busTimes, 0)
		} else {
			time, _ := strconv.Atoi(t)
			busTimes = append(busTimes, time)
		}

	}

	return startTime, busTimes
}

func ProductOfEarliestDepartingBusAndWaitingTime(start int, busTimes []int) int {
	earliestBus, earliestDeparture := 0, math.MaxInt32
	var departure int
	for _, bus := range busTimes {
		if bus == 0 {
			continue
		}
		departure = bus - (start % bus)
		if departure < earliestDeparture {
			earliestBus = bus
			earliestDeparture = departure
		}
	}
	return earliestBus * earliestDeparture
}

// FindIdealStart => find start time such that all buses leave in same time intervals as the offset in busTimes array
func FindIdealStart(busTimes []int) int {
	var start, m, residue int
	// using chinese remainder theorem
	moduloProduct := product(busTimes)
	for i, bus := range busTimes {
		if bus == 0 {
			continue
		}
		m = moduloProduct / bus
		residue = positiveResidue(-i, bus)
		_, inverse, _ := Egcd(m, bus)
		if inverse < 0 {
			inverse = positiveResidue(inverse, bus)
		}
		start += (residue) * m * inverse
	}
	return start % moduloProduct
}

func positiveResidue(n int, modulus int) int {
	for n < 0 {
		n = modulus - (n * -1)
		n = n % modulus
	}
	return n
}

func product(numbers []int) int {
	p := 1
	for _, n := range numbers {
		if n == 0 {
			continue
		}
		p *= n
	}
	return p
}

// Egcd => calculates gcd as well as x, y which satisy a.x + b.y = gcd(a,b)
func Egcd(a int, b int) (int, int, int) {
	x, y := 1, 0
	x1, y1 := 0, 1
	for b != 0 {
		q, r := a/b, a-(a/b)*b
		x, x1 = x1, x-q*x1
		y, y1 = y1, y-q*y1
		a, b = b, r
	}
	return a, x, y
}
