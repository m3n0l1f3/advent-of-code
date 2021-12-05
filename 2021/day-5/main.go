package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func parseCoordinate(input string) coordinate {
	tokens := strings.Split(input, ",")
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return coordinate{x, y}
}

func checkCoordinateLine(start coordinate, end coordinate) bool {
	return start.x == end.x || start.y == end.y
}

func updateCoordinateMaps(maps map[coordinate]int, start coordinate, end coordinate) map[coordinate]int {

	dx := 0
	if start.x > end.x {
		dx = -1
	} else if start.x < end.x {
		dx = 1
	}

	dy := 0
	if start.y > end.y {
		dy = -1
	} else if start.y < end.y {
		dy = 1
	}

	current := start
	for {

		maps[current]++

		if current.x == end.x && current.y == end.y {
			break
		}

		current.x += dx
		current.y += dy
	}
	return maps
}

func partOne(scanner bufio.Scanner) int {

	var coordinateMaps = make(map[coordinate]int)

	for scanner.Scan() {
		lineText := scanner.Text()
		coordinateTokens := strings.Split(lineText, " -> ")
		start := parseCoordinate(coordinateTokens[0])
		end := parseCoordinate(coordinateTokens[1])

		if checkCoordinateLine(start, end) {
			coordinateMaps = updateCoordinateMaps(coordinateMaps, start, end)
		}
	}

	result := 0

	for _, count := range coordinateMaps {
		if count > 1 {
			result++
		}
	}

	return result
}

func partTwo(scanner bufio.Scanner) int {

	var coordinateMaps = make(map[coordinate]int)

	for scanner.Scan() {
		lineText := scanner.Text()
		coordinateTokens := strings.Split(lineText, " -> ")
		start := parseCoordinate(coordinateTokens[0])
		end := parseCoordinate(coordinateTokens[1])

		coordinateMaps = updateCoordinateMaps(coordinateMaps, start, end)
	}

	result := 0

	for _, count := range coordinateMaps {
		if count > 1 {
			result++
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
