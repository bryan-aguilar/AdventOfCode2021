package daytwo

import (
	"testing"
)

var input string = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPartOne(t *testing.T) {

	got := PartOne(input)
	want := 150

	if got != want {
		t.Errorf("Part one failed: got %d, want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	want := 900

	if got != want {
		t.Errorf("Part two failed: got %d, want %d", got, want)
	}
}
