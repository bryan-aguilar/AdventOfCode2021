package dayOne

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func PartOne() (ans int) {
	data, err := os.ReadFile("inputs/dayOneInput.txt")
	check(err)

	dataStr := string(data)
	dataSlc := strings.Split(dataStr, "\n")
	prev, err := strconv.Atoi(dataSlc[0])
	check(err)

	for i := 1; i < len(dataSlc); i++ {
		curval, err := strconv.Atoi(dataSlc[i])
		check(err)
		if curval > prev {
			ans++
		}
		prev = curval
	}
	return
}

func PartTwo() (ans int) {
	data, err := os.ReadFile("inputs/dayOneInput.txt")
	check(err)

	dataStr := string(data)
	dataSlc := strings.Split(dataStr, "\n")
	for i := 3; i < len(dataSlc); i++ {
		prev, err := strconv.Atoi(dataSlc[i-3])
		check(err)
		curval, err := strconv.Atoi(dataSlc[i])
		check(err)
		if curval > prev {
			ans++
		}
	}
	return
}
