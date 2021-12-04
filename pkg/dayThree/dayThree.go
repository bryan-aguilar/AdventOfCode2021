package dayThree

import (
	ec "aoc21/internal/errCheck"
	"strconv"
	"strings"
)

/*
Oh boy day three and I get to work with binary numbers already
The O(n2) solution seems fairly straight forward. We just need to find the most used
common value.
Finding epsilon rate can be done with an XOR
Is there a math library to work with binary values in golang?
*/
func PartOne(input string) (ans int) {
	inputSlc := strings.Split(input, "\n")
	var gamma string
	//figure out gamma value
	for col := 0; col < len(inputSlc[0]); col++ {
		var counter int
		for _, word := range inputSlc {
			if string(word[col]) == "1" {
				counter++
			}
		}
		if counter > (len(inputSlc) / 2) {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	gammaIntVal, err := strconv.ParseInt(gamma, 2, len(inputSlc[0])+1)
	ec.Check(err)

	// create our epsilon value by hand because I was struggling to perform the bit operations on
	// arbitrary bit lengths.
	var epsilon string
	for _, char := range gamma {
		if string(char) == "1" {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}

	epsilonIntVal, err := strconv.ParseInt(epsilon, 2, len(inputSlc[0])+1)

	ans = int(gammaIntVal * epsilonIntVal)
	return
}

func PartTwo(input string) (ans int) {
	inputSlc := strings.Split(input, "\n")

	var c02list []string = make([]string, len(inputSlc))
	var o2list []string = make([]string, len(inputSlc))
	copy(c02list, inputSlc)
	copy(o2list, inputSlc)

	var co2idx int
	//find c02 rating
	for len(c02list) > 1 {
		lc := getLeastCommonValue(c02list, co2idx)
		for i := 0; i < len(c02list); i++ {
			val := c02list[i]
			if string(val[co2idx]) != lc {
				c02list = removeIndex(c02list, i)
				i--
			}
		}
		co2idx++
	}

	//find oxygen rating
	var o2idx int
	for len(o2list) > 1 {
		mc := getMostCommonValue(o2list, o2idx)
		for i := 0; i < len(o2list); i++ {
			val := o2list[i]
			if string(val[o2idx]) != mc {
				o2list = removeIndex(o2list, i)
				i--
			}
		}
		o2idx++
	}

	o2IntVal, err := strconv.ParseInt(o2list[0], 2, len(o2list[0])+1)
	ec.Check(err)

	co2IntVal, err := strconv.ParseInt(c02list[0], 2, len(c02list[0])+1)

	ans = int(o2IntVal * co2IntVal)
	return
}

// expensive but easiest solution atm
func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func getMostCommonValue(input []string, desiredIndex int) string {
	var counter int
	for _, val := range input {
		if string(val[desiredIndex]) == "1" {
			counter++
		}
	}
	// determine if most common is 1 or 0
	if counter >= (len(input) - counter) {
		return "1"
	} else {
		return "0"
	}

}

func getLeastCommonValue(input []string, desiredIndex int) string {
	var counter int
	for _, val := range input {
		if string(val[desiredIndex]) == "1" {
			counter++
		}
	}
	// determine if most common is 1 or 0
	if counter >= (len(input) - counter) {
		return "0"
	} else {
		return "1"
	}

}
