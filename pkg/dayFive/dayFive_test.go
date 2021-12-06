package dayFive

import "testing"

var input string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	want := 5

	if got != want {
		t.Errorf("Day 5 Part one failed got: %d want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	want := 12

	if got != want {
		t.Errorf("Day 5 part two failed got: %d want %d", got, want)
	}
}
