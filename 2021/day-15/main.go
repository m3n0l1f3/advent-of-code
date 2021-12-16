package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func loadData(scanner bufio.Scanner) [][]int {

	maze := make([][]int, 0)

	for scanner.Scan() {
		row := make([]int, 0)
		lineText := scanner.Text()
		for _, r := range lineText {
			s := string(r)
			i, _ := strconv.Atoi(s)
			row = append(row, i)
		}
		maze = append(maze, row)
	}

	return maze
}

func intMin(left int, right int) int {
	if left < right {
		return left
	}
	return right
}

func inplaceSolve(maze [][]int) [][]int {
	rowMax, colMax := len(maze), len(maze[0])

	for i := 0; i < rowMax; i++ {
		for j := 0; j < colMax; j++ {
			if i > 0 && j > 0 {
				maze[i][j] += intMin(maze[i-1][j], maze[i][j-1])
			} else if i > 0 {
				maze[i][j] += maze[i-1][j]
			} else if j > 0 {
				maze[i][j] += maze[i][j-1]
			}
		}
	}

	return maze
}

func partOne(maze [][]int) int {

	maze = inplaceSolve(maze)
	rowMax, colMax := len(maze), len(maze[0])

	return maze[rowMax-1][colMax-1] - maze[0][0]
}

func expand(snippet [][]int, scale int) [][]int {
	fullMap := make([][]int, 0)

	for i := 0; i < scale; i++ {
		for _, row := range snippet {
			fullRow := make([]int, 0)
			for j := 0; j < scale; j++ {
				for _, value := range row {
					fullValue := value + i + j
					if fullValue > 9 {
						fullValue -= 9
					}
					fullRow = append(fullRow, fullValue)
				}
			}
			fullMap = append(fullMap, fullRow)
		}
	}

	return fullMap
}

func partTwo(mazeSnippet [][]int) int {

	maze := expand(mazeSnippet, 5)

	maze = inplaceSolve(maze)
	rowMax, colMax := len(maze), len(maze[0])

	return maze[rowMax-1][colMax-1] - maze[0][0]
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
	maze := loadData(*scanner)
	result := partOne(maze)

	fmt.Println(result)
}
