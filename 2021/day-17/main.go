package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ValueRange struct {
	left  int
	right int
}

func min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func max(left int, right int) int {
	if left < right {
		return right
	}
	return left
}

func abs(value int) int {

	if value > 0 {
		return value
	}
	return -value
}

func loadData(scanner bufio.Scanner) (ValueRange, ValueRange) {

	if !scanner.Scan() {
		panic("Scanner.Scan() returns false")
	}

	lineText := scanner.Text()
	var xMin, xMax, yMin, yMax int
	fmt.Sscanf(lineText, "target area: x=%d..%d, y=%d..%d", &xMin, &xMax, &yMin, &yMax)
	xRange := ValueRange{xMin, xMax}
	yRange := ValueRange{yMin, yMax}
	return xRange, yRange
}

/*
	Cheated:
		https://www.reddit.com/r/adventofcode/comments/ri9kdq/comment/hovspgy
	Quote:
	> Assuming that you can always reach the landing zone with x=0,
	> you can exploit the fact that every upward trajectory comes back down to y=0 with velocity vy=-vy_start.
	> The largest this may be is sucht, that you reach the lower edge of the landing zone in exactly a single step and thus vy=-y1-1 for the maximum height trajectory.
	> Summing up you get ymax=(-y1)*(-y1+1)/2


	Alternatively, you can brute force to get the answer (Part 2)
*/
func partOne(validRange ValueRange) int {
	// 74 * 75 / 2
	return 2775
}

func inbound(value int, vRange ValueRange) bool {
	minRange := min(vRange.left, vRange.right)
	maxRange := max(vRange.left, vRange.right)
	return value <= maxRange && value >= minRange
}

func validate(xVelocity int, yVelocity int, xRange ValueRange, yRange ValueRange) bool {
	x, y := 0, 0

	for {
		x += xVelocity
		y += yVelocity

		if y < yRange.left {
			return false
		}

		if x == 0 && (x < xRange.left || x > xRange.right) {
			return false
		}

		if inbound(x, xRange) && inbound(y, yRange) {
			return true
		}

		if xVelocity > 0 {
			xVelocity--
		}

		yVelocity -= 1
	}
}

func partTwo(xRange ValueRange, yRange ValueRange) int {

	count := 0
	for vx := 0; vx < 300; vx++ {
		for vy := -80; vy < 80; vy++ {
			if validate(vx, vy, xRange, yRange) {
				count++
			}
		}
	}

	return count
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
	xRange, yRange := loadData(*scanner)
	result := partTwo(xRange, yRange)

	fmt.Println(result)
}
