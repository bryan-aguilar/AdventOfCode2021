package daySix

import "testing"

var input string = `3,4,3,1,2`

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	want := 5934

	if got != want {
		t.Errorf("Part one got: %d want: %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input, 256)
	want := 26984457539

	if got != want {
		t.Errorf("Part two got: %d want: %d", got, want)
	}
}
