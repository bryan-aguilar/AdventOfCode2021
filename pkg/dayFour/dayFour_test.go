package dayFour

import (
	"fmt"
	"testing"
)

var input string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	want := 4512

	if got != want {
		t.Errorf("Day Three part one failed got: %d want: %d", got, want)
	}
}

func TestMarkBoard(t *testing.T) {
	mockInput :=
		`22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`
	got := MarkBoard(mockInput, 3, 3)
	want :=
		`22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10 3 * 5
1 12 20 15 19`

	if got != want {
		fmt.Print(got)
		t.Errorf("Mark board failed")
	}

}

func TestCheckWinner(t *testing.T) {
	testColInput :=
		`* 1 1 1 1
*  2 23  4 24
*  9 14 16  7
* 10  3 18  5
* 12 20 15 19`
	got := CheckBoardForWinner(testColInput, 0, 0)
	want := true

	if got != want {
		fmt.Print(got)
		t.Errorf("Check winner failed col check")
	}

	testRowInput :=
		`22 13 17 11  0
8  2 23  4 24
*  * * *  *
6 10  3 18  5
1 12 20 15 19`

	got = CheckBoardForWinner(testRowInput, 2, 2)
	want = true

	if got != want {
		fmt.Print(got)
		t.Errorf("Check winner failed row check")
	}
}

func TestCalculateWinningBoard(t *testing.T) {
	testRealBoard :=
		`0 0 0 0 0
1 2 0 0 0 
0 0 0 0 0 
0 0 0 0 0
1 2 3 4 5`

	testMockBoard :=
		`0 0 0 0 0 
0 0 0 0 0 
0 0 0 0 0 
0 0 0 0 0
* * * * *`

	got := CalculateWinningValue(testRealBoard, testMockBoard, 3)
	want := 9

	if got != want {
		t.Errorf("Calculate winner failed got : %d want: %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	want := 1924

	if got != want {
		t.Errorf("Day 4 Part Two failed got: %d want: %d", got, want)
	}
}
