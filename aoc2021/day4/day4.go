package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	bingo  bool
	values [25]int
	marked [25]bool
}

func readInput(path string) (nums []int, boards []Board, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for _, s := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}

	boardIndex := 0
	var currentBoard Board

	for scanner.Scan() {
		if len(scanner.Text()) > 1 {
			if boardIndex > 0 && boardIndex%5 == 0 {
				boards = append(boards, currentBoard)
			}

			idx := ((boardIndex - (boardIndex % 5)) + (5 * boardIndex)) % 25

			fmt.Sscanf(scanner.Text(), "%d %d %d %d %d",
				&currentBoard.values[idx+0],
				&currentBoard.values[idx+1],
				&currentBoard.values[idx+2],
				&currentBoard.values[idx+3],
				&currentBoard.values[idx+4])

			boardIndex++
		}
	}

	boards = append(boards, currentBoard)

	return nums, boards, nil
}

func checkBingo(checks [25]bool) bool {
	var lines [5]bool = [5]bool{true, true, true, true, true}

	for i := 0; i < 5; i++ {
		lines[0] = lines[0] && checks[i]
		lines[1] = lines[1] && checks[i+5]
		lines[2] = lines[2] && checks[i+10]
		lines[3] = lines[3] && checks[i+15]
		lines[4] = lines[4] && checks[i+20]

		if checks[i] && checks[i+5] && checks[i+10] && checks[i+15] && checks[i+20] {
			return true
		}
	}

	if lines[0] || lines[1] || lines[2] || lines[3] || lines[4] {
		return true
	}

	return false
}

func resetBoard(boards []Board) {
	for i := range boards {
		boards[i].bingo = false
		for b := range boards[i].marked {
			boards[i].marked[b] = false
		}
	}
}

func bingo(numbers []int, boards []Board, winner bool) (board int, lastCall int) {
	var winners int = 0
	var numIdx int = 0
	var numBoards = len(boards)
	var lastWin int

	for winners < numBoards {
		bingoCall := numbers[numIdx : numIdx+5]

		for _, call := range bingoCall {
			for boardIndex := range boards {
				if !boards[boardIndex].bingo {
					for valueIndex := range boards[boardIndex].values {
						if call == boards[boardIndex].values[valueIndex] {
							boards[boardIndex].marked[valueIndex] = true

							if checkBingo(boards[boardIndex].marked) {
								if winner {
									return boardIndex, call
								}

								boards[boardIndex].bingo = true
								lastWin = boardIndex
								lastCall = call
								winners++
							}
						}
					}
				}
			}
		}

		numIdx += 5
	}

	return lastWin, lastCall
}

func sumUnmarkedNumbers(board Board) (result int) {
	for idx, num := range board.values {
		if !board.marked[idx] {
			result += num
		}
	}
	return
}

func Execute() {
	numbers, boards, err := readInput("input/4/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	winningBoard, lastCall := bingo(numbers, boards, true)
	fmt.Printf("Winning Board Sum: %d\n", sumUnmarkedNumbers(boards[winningBoard])*lastCall)

	resetBoard(boards)

	winningBoard, lastCall = bingo(numbers, boards, false)
	fmt.Printf("Losing Board Sum: %d\n\n", sumUnmarkedNumbers(boards[winningBoard])*lastCall)
}
