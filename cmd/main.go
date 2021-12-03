package main

// Take in command line arguments to process which day we want
// Daily problems will be stored in there applicable package directory
import (
	dayOne "aoc21/pkg/dayOne"
	"fmt"
)

func main() {
	fmt.Println(dayOne.PartOne())
	fmt.Println(dayOne.PartTwo())
}
