package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

/*
	[i - 1, j - 1] [i - 1, j] [i - 1, j + 1]
	[i    , j - 1]  Current   [i    , j + 1]
	[i + 1, j - 1] [i + 1, j] [i + 1, j + 1]
*/
func updateGrid(grid [][]int, i int, j int) [][]int {
	rowLimit, colLimit := len(grid), len(grid[0])
	if i > 0 {
		if j > 0 {
			grid[i-1][j-1]++
		}
		grid[i-1][j]++
		if j < colLimit-1 {
			grid[i-1][j+1]++
		}
	}

	if j > 0 {
		grid[i][j-1]++
	}

	if j < colLimit-1 {
		grid[i][j+1]++
	}

	if i < rowLimit-1 {
		if j > 0 {
			grid[i+1][j-1]++
		}
		grid[i+1][j]++
		if j < colLimit-1 {
			grid[i+1][j+1]++
		}
	}

	return grid
}

func flashGrid(grid [][]int) (output [][]int, count int) {

	flashedMap := make(map[coordinate]bool)
	flashList := make([]coordinate, 0)
	for i, row := range grid {
		for j := range row {
			grid[i][j]++
			if grid[i][j] > 9 {
				flashList = append(flashList, coordinate{i, j})
			}
		}
	}

	for len(flashList) > 0 {
		newFlashList := make([]coordinate, 0)
		for _, coords := range flashList {

			if flashedMap[coords] {
				continue
			}

			flashedMap[coords] = true
			grid = updateGrid(grid, coords.x, coords.y)

			for i, row := range grid {
				for j, value := range row {
					currentCoords := coordinate{i, j}
					if value > 9 && !flashedMap[currentCoords] {
						newFlashList = append(newFlashList, currentCoords)
					}
				}
			}
		}
		flashList = newFlashList
	}

	for i, row := range grid {
		for j, value := range row {
			if value > 9 {
				grid[i][j] = 0
				count++
			}
		}
	}

	output = grid
	return
}

func partOne(scanner bufio.Scanner) int {

	grid := make([][]int, 0)
	for scanner.Scan() {
		lineText := scanner.Text()
		row := make([]int, 0)
		for _, r := range lineText {
			digit, _ := strconv.Atoi(string(r))
			row = append(row, digit)
		}
		grid = append(grid, row)
	}

	result := 0
	flashCount := 0
	for i := 0; i < 100; i++ {
		grid, flashCount = flashGrid(grid)
		fmt.Println("[", i, "]", flashCount)
		for _, row := range grid {
			fmt.Println(row)
		}
		fmt.Println()
		result += flashCount
	}

	return result
}

func partTwo(scanner bufio.Scanner) int {

	gridSize := 0
	grid := make([][]int, 0)
	for scanner.Scan() {
		lineText := scanner.Text()
		row := make([]int, 0)
		for _, r := range lineText {
			digit, _ := strconv.Atoi(string(r))
			row = append(row, digit)
			gridSize++
		}
		grid = append(grid, row)
	}

	result := 0
	flashCount := 0
	for {
		result++
		grid, flashCount = flashGrid(grid)

		if flashCount == gridSize {
			break
		}
	}
	return result
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
	result := partTwo(*scanner)

	fmt.Println(result)
}
