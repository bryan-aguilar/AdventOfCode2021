package dayFive

import (
	ec "aoc21/internal/errCheck"
	"sort"
	"strconv"
	"strings"
)

/*
	- General Strategy at a first glance
	- We just keep track of of how many times a specific x,y coordinate has been hit
	- This would mean we need to make some type of map that maps (x,y) -> int
	- First glance...not sure how I may this map. Maybe a struct that acts like a point?
	- bit of research, a point struct seems to be the answer

*/
type Point struct {
	x int
	y int
}

func PartOne(rawInput string) (ans int) {
	// split on new line for each input
	input := strings.Split(rawInput, "\n")
	hitMap := make(map[Point]int)
	for _, inputVal := range input {
		// split our values
		tmpSlic := strings.Split(inputVal, " -> ")
		//pull out start and end
		startValStr := tmpSlic[0]
		endValStr := tmpSlic[1]

		//create point structs to reflect start and end
		startPt := parsePoint(startValStr)
		endPt := parsePoint(endValStr)

		// if x values are same then we are going horizontal
		// if y values are the same then we are going vertical
		if startPt.x == endPt.x {
			// simple slice to establish start and end
			startEnd := []int{startPt.y, endPt.y}
			sort.Ints(startEnd)
			for curYval := startEnd[0]; curYval <= startEnd[1]; curYval++ {
				//create Point for current position
				curPoint := Point{startPt.x, curYval}
				hitPoint(hitMap, curPoint)
			}
		} else if startPt.y == endPt.y {
			startEnd := []int{startPt.x, endPt.x}
			sort.Ints(startEnd)
			for curXval := startEnd[0]; curXval <= startEnd[1]; curXval++ {
				curPoint := Point{curXval, startPt.y}
				hitPoint(hitMap, curPoint)
			}
		}
	}

	ans = finalScoreCalc(hitMap)
	return
}

/*
	- Dumb brute force solution should go as follows
	- for any diagonal lines choose a point in the pair as the starting point
	- from there we can only move down left, down right, up left, or up right
	- figure that out and do loops
*/
func PartTwo(rawInput string) (ans int) {
	// split on new line for each input
	input := strings.Split(rawInput, "\n")
	hitMap := make(map[Point]int)
	for _, inputVal := range input {
		// split our values
		tmpSlic := strings.Split(inputVal, " -> ")
		//pull out start and end
		startValStr := tmpSlic[0]
		endValStr := tmpSlic[1]

		//create point structs to reflect start and end
		startPt := parsePoint(startValStr)
		endPt := parsePoint(endValStr)

		// if x values are same then we are going horizontal
		// if y values are the same then we are going vertical
		if startPt.x == endPt.x {
			// simple slice to establish start and end
			startEnd := []int{startPt.y, endPt.y}
			sort.Ints(startEnd)
			for curYval := startEnd[0]; curYval <= startEnd[1]; curYval++ {
				//create Point for current position
				curPoint := Point{startPt.x, curYval}
				hitPoint(hitMap, curPoint)
			}
		} else if startPt.y == endPt.y {
			startEnd := []int{startPt.x, endPt.x}
			sort.Ints(startEnd)
			for curXval := startEnd[0]; curXval <= startEnd[1]; curXval++ {
				curPoint := Point{curXval, startPt.y}
				hitPoint(hitMap, curPoint)
			}
		} else {
			// diagonals
			// our logic below requires us to hit the start point manually
			curPoint := startPt
			hitPoint(hitMap, startPt)
			for curPoint != endPt {
				// end is up and to the right
				if endPt.x > startPt.x && endPt.y < startPt.y {
					curPoint.x++
					curPoint.y--
					hitPoint(hitMap, curPoint)
					// up and to the left
				} else if endPt.x < startPt.x && endPt.y < startPt.y {
					curPoint.x--
					curPoint.y--
					hitPoint(hitMap, curPoint)
					// down and to the right
				} else if endPt.x > startPt.x && endPt.y > startPt.y {
					curPoint.x++
					curPoint.y++
					hitPoint(hitMap, curPoint)
					// down and to the left
				} else if endPt.x < startPt.x && endPt.y > startPt.y {
					curPoint.x--
					curPoint.y++
					hitPoint(hitMap, curPoint)
				}
			}
		}
	}
	// sum up all values in map that are >= 2
	ans = finalScoreCalc(hitMap)
	return
}

func finalScoreCalc(m map[Point]int) (ans int) {
	for _, hits := range m {
		if hits >= 2 {
			ans++
		}
	}
	return
}

// incrememnts our map for us
func hitPoint(m map[Point]int, p Point) {
	if curCounter, exists := m[p]; exists {
		m[p] = curCounter + 1
	} else {
		m[p] = 1
	}
}

// expect "4,2" return Point{4,2}
func parsePoint(rawPt string) Point {
	tmp := strings.Split(rawPt, ",")
	x, err := strconv.Atoi(tmp[0])
	ec.Check(err)
	y, err := strconv.Atoi(tmp[1])
	ec.Check(err)
	return Point{x, y}
}
