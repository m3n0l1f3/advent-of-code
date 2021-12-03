package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(values []int) int {
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

func partOne(scanner bufio.Scanner) (count int) {
	// Assume input numbers are always above 0
	currentValue := 0
	count = 0

	for scanner.Scan() {
		nextValue, _ := strconv.Atoi(scanner.Text())
		if currentValue != 0 && nextValue > currentValue {
			count += 1
		}

		currentValue = nextValue
	}

	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
		os.Exit(1)
	}

	return
}

func partTwo(scanner bufio.Scanner, windowSize int) (count int) {
	if windowSize == 0 {
		windowSize = 3
	}

	currentWindow := make([]int, 0, windowSize)
	currentSum := 0
	count = 0

	for scanner.Scan() {
		nextValue, _ := strconv.Atoi(scanner.Text())
		if len(currentWindow) < windowSize {
			currentWindow = append(currentWindow, nextValue)
			continue
		}
		currentSum = sum(currentWindow)
		nextSum := currentSum - currentWindow[0] + nextValue
		if nextSum > currentSum {
			count += 1
		}
		currentWindow = currentWindow[1:windowSize]
		currentWindow = append(currentWindow, nextValue)
	}

	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
		os.Exit(1)
	}

	return
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

	// count := partOne(*scanner)
	count := partTwo(*scanner, 3)

	fmt.Println(count)
}
