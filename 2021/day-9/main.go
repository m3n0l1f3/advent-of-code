package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func isLowestPoint(heightMap [][]int, rowIndex int, colIndex int) bool {
	rowLimit, colLimit := len(heightMap), len(heightMap[0])

	upRow, upCol := rowIndex-1, colIndex
	leftRow, leftCol := rowIndex, colIndex-1
	rightRow, rightCol := rowIndex, colIndex+1
	downRow, downCol := rowIndex+1, colIndex
	currentHeight := heightMap[rowIndex][colIndex]

	if upRow >= 0 && heightMap[upRow][upCol] <= currentHeight {
		return false
	}

	if leftCol >= 0 && heightMap[leftRow][leftCol] <= currentHeight {
		return false
	}

	if rightCol < colLimit && heightMap[rightRow][rightCol] <= currentHeight {
		return false
	}

	if downRow < rowLimit && heightMap[downRow][downCol] <= currentHeight {
		return false
	}

	return true
}

func partOne(scanner bufio.Scanner) int {

	var heightMap = make([][]int, 0)

	for scanner.Scan() {
		lineText := scanner.Text()
		var heightRow = make([]int, 0)
		for _, r := range lineText {
			height, _ := strconv.Atoi(string(r))
			heightRow = append(heightRow, height)
		}
		heightMap = append(heightMap, heightRow)
	}

	result := 0
	for i, row := range heightMap {
		for j, height := range row {
			if isLowestPoint(heightMap, i, j) {
				result += height + 1
			}
		}
	}

	return result
}

func dfs(visitedMap *[][]bool, heightMap *[][]int, rowIndex int, colIndex int) int {
	rowLimit, colLimit := len(*heightMap), len((*heightMap)[0])

	if rowIndex < 0 || colIndex < 0 || rowIndex >= rowLimit || colIndex >= colLimit {
		return 0
	}
	if (*visitedMap)[rowIndex][colIndex] {
		return 0
	}

	(*visitedMap)[rowIndex][colIndex] = true
	if (*heightMap)[rowIndex][colIndex] == 9 {
		return 0
	}

	return 1 +
		dfs(visitedMap, heightMap, rowIndex-1, colIndex) +
		dfs(visitedMap, heightMap, rowIndex+1, colIndex) +
		dfs(visitedMap, heightMap, rowIndex, colIndex-1) +
		dfs(visitedMap, heightMap, rowIndex, colIndex+1)
}

func partTwo(scanner bufio.Scanner) int {

	var heightMap = make([][]int, 0)
	var visitedMap = make([][]bool, 0)
	for scanner.Scan() {
		lineText := scanner.Text()
		var heightRow = make([]int, 0)
		var visitedRow = make([]bool, 0)
		for _, r := range lineText {
			height, _ := strconv.Atoi(string(r))
			heightRow = append(heightRow, height)
			if height == 9 {
				visitedRow = append(visitedRow, true)
			} else {
				visitedRow = append(visitedRow, false)
			}
		}
		visitedMap = append(visitedMap, visitedRow)
		heightMap = append(heightMap, heightRow)
	}

	sizeList := make([]int, 0)
	for i, row := range heightMap {
		for j := range row {
			if !visitedMap[i][j] {
				sizeList = append(sizeList, dfs(&visitedMap, &heightMap, i, j))
			}
		}
	}

	sort.Ints(sizeList)

	totalLeng := len(sizeList)

	return sizeList[totalLeng-1] * sizeList[totalLeng-2] * sizeList[totalLeng-3]
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
