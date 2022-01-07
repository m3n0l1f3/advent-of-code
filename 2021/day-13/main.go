package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dot struct {
	x int
	y int
}

func parse(scanner bufio.Scanner) (map[Dot]bool, *bufio.Scanner) {
	dotMap := make(map[Dot]bool)
	for {
		scanner.Scan()
		lineText := scanner.Text()
		if len(lineText) == 0 {
			break
		}
		lineToken := strings.Split(lineText, ",")
		xToken, yToken := lineToken[0], lineToken[1]
		x, _ := strconv.Atoi(xToken)
		y, _ := strconv.Atoi(yToken)

		xyDot := Dot{x, y}
		dotMap[xyDot] = true
	}

	return dotMap, &scanner
}

func getDirection(line string) (string, int) {
	lineToken := strings.Split(line, " ")
	directionText := lineToken[2]
	directionToken := strings.Split(directionText, "=")
	direction, marginText := directionToken[0], directionToken[1]
	margin, _ := strconv.Atoi(marginText)

	return direction, margin
}

func fold(dotMap map[Dot]bool, direction string, margin int) map[Dot]bool {
	newDotMap := make(map[Dot]bool)

	for dot := range dotMap {
		if direction == "x" {
			// Fold along x

			if dot.x < margin {
				newDotMap[dot] = true
				continue
			}

			newX := 2*margin - dot.x
			newDot := Dot{newX, dot.y}
			newDotMap[newDot] = true

		} else {
			// Fold along y
			if dot.y < margin {
				newDotMap[dot] = true
				continue
			}

			newY := 2*margin - dot.y
			newDot := Dot{dot.x, newY}
			newDotMap[newDot] = true
		}
	}

	return newDotMap
}

func partOne(scanner bufio.Scanner, dotMap map[Dot]bool) int {

	for scanner.Scan() {
		lineText := scanner.Text()
		direction, margin := getDirection(lineText)
		dotMap = fold(dotMap, direction, margin)
		fmt.Println(len(dotMap))
	}

	return len(dotMap)
}

func intMax(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

func visualize(dotMap map[Dot]bool) {
	maxX, maxY := 0, 0
	for dot := range dotMap {
		maxX = intMax(maxX, dot.x)
		maxY = intMax(maxY, dot.y)
	}

	visualDotMap := make([][]string, 0)

	for i := 0; i <= maxY; i++ {
		visualDotRow := make([]string, 0)
		for j := 0; j <= maxX; j++ {
			dot := Dot{j, i}
			if dotMap[dot] {
				visualDotRow = append(visualDotRow, "#")
			} else {
				visualDotRow = append(visualDotRow, ".")
			}
		}
		visualDotMap = append(visualDotMap, visualDotRow)
	}

	for _, row := range visualDotMap {
		fmt.Println(row)
	}
}

func partTwo(scanner bufio.Scanner, dotMap map[Dot]bool) int {

	for scanner.Scan() {
		lineText := scanner.Text()
		direction, margin := getDirection(lineText)
		dotMap = fold(dotMap, direction, margin)
		visualize(dotMap)
		fmt.Println()
	}

	return len(dotMap)
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

	var dotMap map[Dot]bool
	dotMap, scanner = parse(*scanner)

	result := partTwo(*scanner, dotMap)

	fmt.Println(result)
}
