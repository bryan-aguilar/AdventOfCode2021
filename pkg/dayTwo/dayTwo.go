package daytwo

import (
	ec "aoc21/internal/errCheck"
	"os"
	"strconv"
	"strings"
)

func PartOne() (ans int) {
	data, err := os.ReadFile("inputs/dayTwoInput.txt")
	ec.Check(err)

	dataStr := strings.Split(string(data), "\n")

	var depth, distance int
	for _, val := range dataStr {
		x := strings.Split(val, " ")
		valNum, err := strconv.Atoi(x[1])
		ec.Check(err)
		switch x[0] {
		case "forward":
			distance += valNum
		case "up":
			depth -= valNum
		case "down":
			depth += valNum
		}
	}
	ans = depth * distance
	return
}

func PartTwo() (ans int) {
	data, err := os.ReadFile("inputs/dayTwoInput.txt")
	ec.Check(err)

	dataStr := strings.Split(string(data), "\n")

	var depth, distance, aim int
	for _, val := range dataStr {
		x := strings.Split(val, " ")
		valNum, err := strconv.Atoi(x[1])
		ec.Check(err)
		switch x[0] {
		case "forward":
			distance += valNum
			depth += (aim * valNum)
		case "up":
			aim -= valNum
		case "down":
			aim += valNum
		}
	}
	ans = depth * distance
	return
}
