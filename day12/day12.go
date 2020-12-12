package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Navigation struct {
	direction byte
	value     float64
}

type Point struct {
	X float64
	Y float64
}

// SeatRule => function which takes in seatingMatrix and row, col of seat and returns numnber of occupied seats
//type SeatRule func([][]rune, int, int) int

func ParseDirections(path string) []Navigation {
	filename := fmt.Sprintf("%s/day12/day12.txt", path)
	//filename := fmt.Sprintf("%s/day12/test.txt", path)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	var navigations []Navigation
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(string(line[1:]))
		nav := Navigation{direction: line[0], value: float64(val)}
		navigations = append(navigations, nav)
	}
	return navigations
}

func FollowDirections(navigations []Navigation) Point {
	final := Point{}
	currentAxis := 0.0 // represented in degrees (currently 0 as we start with ship facing east)
	var sin, cos float64
	var h float64
	for _, navigation := range navigations {
		h = navigation.value
		switch navigation.direction {
		case byte('E'):
			sin, cos = math.Sincos(0)
		case byte('S'):
			sin, cos = math.Sincos(0.5 * math.Pi)
		case byte('W'):
			sin, cos = math.Sincos(math.Pi)
		case byte('N'):
			sin, cos = math.Sincos(1.5 * math.Pi)
		case byte('F'):
			sin, cos = math.Sincos(currentAxis * math.Pi / 180.0)
		case byte('R'):
			currentAxis += navigation.value
			sin, cos = 0.0, 0.0
		case byte('L'):
			currentAxis -= navigation.value
			sin, cos = 0.0, 0.0
		}
		final.X, final.Y = final.X+h*cos, final.Y+h*sin
	}
	return final
}

func FollowDirectionsRelativeToWayPoint(navigations []Navigation, waypoint Point) Point {
	final := Point{}
	var sin, cos float64
	var theta float64 //angle between waypooint and ship in radians
	for _, navigation := range navigations {
		switch navigation.direction {
		case byte('E'):
			sin, cos = math.Sincos(0.0)
			waypoint.X, waypoint.Y = waypoint.X+navigation.value*cos, waypoint.Y+navigation.value*sin
		case byte('S'):
			sin, cos = math.Sincos(0.5 * math.Pi)
			waypoint.X, waypoint.Y = waypoint.X+navigation.value*cos, waypoint.Y+navigation.value*sin
		case byte('W'):
			sin, cos = math.Sincos(math.Pi)
			waypoint.X, waypoint.Y = waypoint.X+navigation.value*cos, waypoint.Y+navigation.value*sin
		case byte('N'):
			sin, cos = math.Sincos(1.5 * math.Pi)
			waypoint.X, waypoint.Y = waypoint.X+navigation.value*cos, waypoint.Y+navigation.value*sin
		case byte('F'):
			final.X, final.Y = final.X+navigation.value*waypoint.X, final.Y+navigation.value*waypoint.Y
		case byte('R'):
			theta = math.Pi * (navigation.value / 180.0) // in radians
			sin, cos = math.Sincos(-theta)               // '-' since its clockwise
			waypoint.X, waypoint.Y = waypoint.X*cos+waypoint.Y*sin, -waypoint.X*sin+waypoint.Y*cos
		case byte('L'):
			theta = math.Pi * (navigation.value / 180.0) // in radians
			sin, cos = math.Sincos(theta)
			waypoint.X, waypoint.Y = waypoint.X*cos+waypoint.Y*sin, -waypoint.X*sin+waypoint.Y*cos
		}
	}
	return final
}

func ManhattanDistance(p1 Point, p2 Point) float64 {
	return math.Abs(p1.X-p2.X) + math.Abs(p1.Y-p2.Y)
}
