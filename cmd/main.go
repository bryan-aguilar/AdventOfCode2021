package main

// Take in command line arguments to process which day we want
// Daily problems will be stored in there applicable package directory
import (
	dayOne "aoc21/pkg/dayOne"
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
	}
	fmt.Printf("Part one: %d \nPart Two: %d \n", p1, p2)
}
