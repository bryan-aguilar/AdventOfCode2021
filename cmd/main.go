package main

// Take in command line arguments to process which day we want
// Daily problems will be stored in there applicable package directory
import (
	"aoc21/pkg/dayFour"
	dayOne "aoc21/pkg/dayOne"
	"aoc21/pkg/dayThree"
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
	case 3:
		input, err := os.ReadFile("inputs/dayThreeInput.txt")
		check(err)
		p1 = dayThree.PartOne(string(input))
		p2 = dayThree.PartTwo(string(input))
	case 4:
		input, err := os.ReadFile("inputs/dayFourInput.txt")
		check(err)
		p1 = dayFour.PartOne(string(input))
		p2 = dayFour.PartTwo(string(input))
	}
	fmt.Printf("Part one: %d \nPart Two: %d \n", p1, p2)
}
