package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type token struct {
	value   int
	visited bool
}

func parseInput(scanner bufio.Scanner) (bingoInput string, boards [][][]token) {
	scanner.Scan()
	bingoInput = scanner.Text()
	boards = make([][][]token, 0)
	var board [][]token
	for scanner.Scan() {
		scanText := scanner.Text()
		if len(scanText) == 0 {
			if len(board) != 0 {
				boards = append(boards, board)
			}
			board = make([][]token, 0)
		} else {
			scanTokens := strings.Split(scanText, " ")
			boardRow := make([]token, 0)
			for _, scanToken := range scanTokens {
				if scanToken == "" {
					continue
				}
				scanInt, _ := strconv.Atoi(scanToken)

				boardRow = append(boardRow, token{scanInt, false})
			}
			board = append(board, boardRow)
		}
	}
	return
}

func checkBoard(board [][]token) bool {
	boardSize := len(board)

	for i := 0; i < boardSize; i++ {
		rowCheckCount := 0
		colCheckCount := 0
		for j := 0; j < boardSize; j++ {
			if board[i][j].visited {
				rowCheckCount++
			}
			if board[j][i].visited {
				colCheckCount++
			}
		}
		if rowCheckCount == boardSize || colCheckCount == boardSize {
			return true
		}
	}
	return false
}

func sumBoard(board [][]token) int {
	result := 0
	for _, row := range board {
		for _, boardToken := range row {
			if !boardToken.visited {
				result += boardToken.value
			}
		}
	}
	return result
}

func updateBoard(board [][]token, value int) bool {
	for i, row := range board {
		for j, boardToken := range row {
			if boardToken.value == value {
				board[i][j].visited = true
				return true
			}
		}
	}
	return false
}

func partOne(bingoInput string, bingoBoards [][][]token) int {
	inputTokens := strings.Split(bingoInput, ",")

	for _, inputToken := range inputTokens {
		inputNumber, _ := strconv.Atoi(inputToken)
		for _, bingoBoard := range bingoBoards {
			didUpdate := updateBoard(bingoBoard, inputNumber)
			if didUpdate {
				didBingo := checkBoard(bingoBoard)
				if didBingo {
					return sumBoard(bingoBoard) * inputNumber
				}
			}
		}
	}

	return 0
}

func partTwo(bingoInput string, bingoBoards [][][]token) int {
	boards := bingoBoards
	inputTokens := strings.Split(bingoInput, ",")
	lastBingoBoard := 0

	for _, inputToken := range inputTokens {
		inputNumber, _ := strconv.Atoi(inputToken)
		newBoards := make([][][]token, 0)
		for _, bingoBoard := range boards {
			didUpdate := updateBoard(bingoBoard, inputNumber)
			if didUpdate {
				didBingo := checkBoard(bingoBoard)
				if didBingo {
					lastBingoBoard = sumBoard(bingoBoard) * inputNumber
					continue
				}
			}

			newBoards = append(newBoards, bingoBoard)
		}
		boards = newBoards
	}

	return lastBingoBoard
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Missing Input File Path")
		os.Exit(1)
	}

	// Open File
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bingoInput, boards := parseInput(*scanner)
	result := partTwo(bingoInput, boards)

	fmt.Println(result)
}
