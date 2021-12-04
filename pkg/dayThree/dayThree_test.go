package dayThree

import "testing"

var input string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	want := 198

	if got != want {
		t.Errorf("Day three part one failed, got %d want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	want := 230

	if got != want {
		t.Errorf("Day three part two failed, got %d want %d", got, want)
	}
}
