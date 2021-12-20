package daySix

import (
	ec "aoc21/internal/errCheck"
	"fmt"
	"strconv"
	"strings"
)

func getIntSlc(input string) []int {
	strSlc := strings.Split(input, ",")
	var outputSlc []int
	for _, val := range strSlc {
		conv, err := strconv.Atoi(val)
		ec.Check(err)
		outputSlc = append(outputSlc, conv)
	}
	return outputSlc
}

func PartOne(input string) int {
	inputSlc := getIntSlc(input)
	numDays := 80
	for numDays > 0 {
		var valuesToAppend []int
		for i, val := range inputSlc {
			if val == 0 {
				valuesToAppend = append(valuesToAppend, 8)
				inputSlc[i] = 6
			} else {
				inputSlc[i]--
			}
		}
		numDays--
		inputSlc = append(inputSlc, valuesToAppend...)
	}
	return len(inputSlc)
}

// part two is a little tricky. We want to avoid the n2 nature of the way we handled part one
// my solution to this will just to use a hashmap and shift values around

func PartTwo(input string, numDays int) (ans int) {
	inputSlc := getIntSlc(input)

	lifeMap := make(map[int]int)

	// populate our initial map with 0 values
	// I don't know if we actually need to do this. We may be able to just use
	// zero values. Will experiment later
	for i := 0; i <= 8; i++ {
		lifeMap[i] = 0
	}

	// map our input onto the map
	for _, val := range inputSlc {
		lifeMap[val]++
	}
	fmt.Println(lifeMap)
	// I believe this would be something like o(n*9)(?) where N is the number of days
	// which is much better our last algo for part one. I think this would technically be o(n) since
	// we drop that constant(?)

	// we will write all changes to this new map and then copy
	newMap := make(map[int]int)
	for numDays > 0 {
		for key, val := range lifeMap {
			// we have special case for 0
			// where we create a new value and then restart at 6
			if key == 0 {
				newMap[8] += val
				newMap[6] += val
			} else {
				newMap[key-1] += val
			}
		}

		for key, val := range newMap {
			lifeMap[key] = val
			newMap[key] = 0
		}
		numDays--
	}

	for _, val := range lifeMap {
		ans += val
	}
	return
}
