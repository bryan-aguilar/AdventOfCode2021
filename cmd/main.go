package main

// Take in command line arguments to process which day we want
// Daily problems will be stored in there applicable package directory
import (
	dayOne "aoc21/pkg/dayOne"
	dayTwo "aoc21/pkg/dayTwo"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dayCmd, err := strconv.Atoi(os.Args[1])
	check(err)
	var p1, p2 int
	switch dayCmd {
	case 1:
		p1 = dayOne.PartOne()
		p2 = dayOne.PartTwo()
	case 2:
		input, err := os.ReadFile("inputs/dayTwoInput.txt")
		check(err)
		p1 = dayTwo.PartOne(string(input))
		p2 = dayTwo.PartTwo(string(input))
	}
	fmt.Printf("Part one: %d \nPart Two: %d \n", p1, p2)
}
