package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseDirection(str string) (horizontal int, vertical int) {
	tokens := strings.Split(str, " ")
	magnitude, _ := strconv.Atoi(tokens[1])

	if tokens[0] == "down" {
		vertical = magnitude
	} else if tokens[0] == "up" {
		vertical = -magnitude
	} else if tokens[0] == "forward" {
		horizontal = magnitude
	}

	return
}

func partOne(scanner bufio.Scanner) int {

	currentHorizontal := 0
	currentVertical := 0

	for scanner.Scan() {
		horizontal, vertical := parseDirection(scanner.Text())
		currentHorizontal += horizontal
		currentVertical += vertical
	}

	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
		os.Exit(1)
	}

	return currentHorizontal * currentVertical
}

func partTwo(scanner bufio.Scanner) int {

	currentHorizontal := 0
	currentVertical := 0
	aim := 0

	for scanner.Scan() {
		horizontal, vertical := parseDirection(scanner.Text())
		currentHorizontal += horizontal
		aim += vertical
		if horizontal > 0 {
			currentVertical += aim * horizontal
		}
	}

	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
		os.Exit(1)
	}

	return currentHorizontal * currentVertical
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
