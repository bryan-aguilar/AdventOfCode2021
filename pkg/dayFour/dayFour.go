package dayFour

import (
	ec "aoc21/internal/errCheck"
	"strconv"
	"strings"
)

/*
	Brute force solution should go as follows:
	- parse and seperate input
	- Store boards and answers seperately
	- make copy of boards
	- If number is selected change char value to "*"
	- check to see if entire row or column is * values
	- if all get output by iterating over board ignoring numbers that have been selected

	// it may be smarter to do this with maps mimicking set behavior.
	// doing some type of math calculation to see if the board has had 5 hits in a specific row or column using a counter
	// then just iterating over the board and ignoring values in the map
*/
func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func SplitInput(input string) (bingoNums []string, bingoBoards []string) {
	firstNewLineIdx := strings.Index(input, "\n")
	withComma := input[:firstNewLineIdx]
	bingoNums = strings.Split(withComma, ",")

	// filter out all of our boards into their own string slice
	split := strings.Split(input, "\n")
	for i := 0; i < len(split); i++ {
		if string(split[i]) == "" {
			split = removeIndex(split, i)
			i--
		}
	}
	split = removeIndex(split, 0)

	for i := 0; i < len(split); i += 5 {
		bingoBoards = append(bingoBoards, strings.Join(split[i:i+5], "\n"))
	}
	return
}

/*
	Recap on part one:
	- From what I read it's not very idiomatic to test unexported functions in golang. This was
	a hard requirement for me because it ruled out anything going wrong in these and causing further issues.
	Unit testing those functions showed me how that splitting the rows using .split() was incorrect because
	it gave me unintended whitespace. .Fields was better once I discovered that.
	- Parsing the input sucked. I wonder if there is a more effecient way to do it.
	- Not having a built in set data type was a little rough. I could have used map but I felt like my method worked.
	- I used a goto. If I wanted to optimize this the first thing I would look at is my looping structure.
*/
func PartOne(input string) (ans int) {
	bingoNums, bingoBoards := SplitInput(input)
	bingoBoardsCopy := make([]string, len(bingoBoards))
	copy(bingoBoardsCopy, bingoBoards)

	for _, bingoNum := range bingoNums {
		for bingoBoardIdx, bingoBoard := range bingoBoards {

			boardRows := strings.Split(bingoBoard, "\n")
			//mark temp board
			for rowIdx, rowVal := range boardRows {
				//split
				rowVals := strings.Fields(rowVal)
				//iterate to see if it contains our value
				for colValsIdx, colVal := range rowVals {
					if string(bingoNum) == string(colVal) {
						bingoBoardsCopy[bingoBoardIdx] = MarkBoard(bingoBoardsCopy[bingoBoardIdx], rowIdx, colValsIdx)
						//determing if board is winner
						if CheckBoardForWinner(bingoBoardsCopy[bingoBoardIdx], rowIdx, colValsIdx) {
							// calculate winning value
							bingoNumInt, err := strconv.Atoi(string(bingoNum))
							ec.Check(err)
							ans = CalculateWinningValue(bingoBoard, bingoBoardsCopy[bingoBoardIdx], bingoNumInt)
							return
						} else {
							// not a winner so go to next board
							// I don't think goto's are a good habit but here we use it to exit the next of for loops
							// this could possibly be avoided by abstracting away the logic above into functions
							goto nextBoard
						}
					}
				}
			}
		nextBoard:
		}
	}
	return
}

/*
 - Brute force solution that is similar to part one
 - Just iterate through nums and keep track of the board state of the last board that won
 - We will need to track board, marking board, and number that triggered
 - we also need to make sure that the board is no longer processed if it has already won
 - calculate winning score at end
*/
func PartTwo(input string) (ans int) {
	bingoNums, bingoBoards := SplitInput(input)
	bingoBoardsCopy := make([]string, len(bingoBoards))
	copy(bingoBoardsCopy, bingoBoards)
	var lastWinningBoard, lastWinningMarkBoard, lastNumberTriggeredWin string
	alreadyWonMap := map[int]bool{}

	for _, bingoNum := range bingoNums {
		for bingoBoardIdx, bingoBoard := range bingoBoards {
			// only process board if it hasnt won
			if _, ok := alreadyWonMap[bingoBoardIdx]; ok {
				continue
			}
			boardRows := strings.Split(bingoBoard, "\n")
			for rowIdx, rowVal := range boardRows {
				rowVals := strings.Fields(rowVal)
				for colValsIdx, colVal := range rowVals {
					if string(bingoNum) == string(colVal) {
						bingoBoardsCopy[bingoBoardIdx] = MarkBoard(bingoBoardsCopy[bingoBoardIdx], rowIdx, colValsIdx)
						if CheckBoardForWinner(bingoBoardsCopy[bingoBoardIdx], rowIdx, colValsIdx) {
							lastWinningBoard = bingoBoard
							lastWinningMarkBoard = bingoBoardsCopy[bingoBoardIdx]
							lastNumberTriggeredWin = bingoNum
							alreadyWonMap[bingoBoardIdx] = true
						} else {
							goto nextBoard
						}
					}
				}
			}
		nextBoard:
		}
	}

	bingoNumInt, err := strconv.Atoi(string(lastNumberTriggeredWin))
	ec.Check(err)
	ans = CalculateWinningValue(lastWinningBoard, lastWinningMarkBoard, bingoNumInt)
	return
}

func CalculateWinningValue(realBoard string, mockBoard string, finalNumCalled int) int {
	// iterate over real board
	realBoardRows := strings.Split(realBoard, "\n")
	mockBoardRows := strings.Split(mockBoard, "\n")

	var boardSum int
	for boardRowIdx, realBoardRow := range realBoardRows {
		realRowVals := strings.Fields(realBoardRow)
		mockRowVals := strings.Fields(mockBoardRows[boardRowIdx])

		for rowColIdx, rowColVal := range realRowVals {
			if string(mockRowVals[rowColIdx]) != string("*") {
				x, err := strconv.Atoi(rowColVal)
				ec.Check(err)
				boardSum += x
			}
		}
	}
	return boardSum * finalNumCalled
}

func CheckBoardForWinner(inputBoard string, rowIdx int, colIdx int) bool {
	boardRows := strings.Split(inputBoard, "\n")
	//check row
	var rowCounter int
	for _, rowVal := range strings.Fields(boardRows[rowIdx]) {
		if string(rowVal) == string("*") {
			rowCounter++
		}
	}
	if rowCounter == 5 {
		return true
	}

	//check cols
	var colCounter int
	for _, boardRow := range boardRows {
		// split row
		rowVals := strings.Fields(boardRow)
		if string(rowVals[colIdx]) == string("*") {
			colCounter++
		}
	}
	if colCounter == 5 {
		return true
	}
	return false
}

func MarkBoard(inputBoard string, rowIdx int, colIdx int) (newBoard string) {
	// split our board into rows
	boardRows := strings.Split(inputBoard, "\n")
	// get our row we want to manipulate
	row := boardRows[rowIdx]
	// split our row
	rowSlice := strings.Fields(row)
	// manipulate
	rowSlice[colIdx] = "*"
	// join back
	manipulatedRow := strings.Join(rowSlice, " ")
	// replace in original
	boardRows[rowIdx] = manipulatedRow
	//join final
	newBoard = strings.Join(boardRows, "\n")
	return
}
