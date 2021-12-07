package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne(scanner bufio.Scanner) int {

	var numberMaps = make(map[int]int)

	for scanner.Scan() {
		lineText := scanner.Text()
		numberStringTokens := strings.Split(lineText, ",")
		for _, numberString := range numberStringTokens {
			number, _ := strconv.Atoi(numberString)
			numberMaps[number]++
		}
	}

	for i := 0; i < 80; i++ {
		newNumberMaps := make(map[int]int)
		newNumberMaps[0] = numberMaps[1]
		newNumberMaps[1] = numberMaps[2]
		newNumberMaps[2] = numberMaps[3]
		newNumberMaps[3] = numberMaps[4]
		newNumberMaps[4] = numberMaps[5]
		newNumberMaps[5] = numberMaps[6]
		newNumberMaps[6] = numberMaps[7] + numberMaps[0]
		newNumberMaps[7] = numberMaps[8]
		newNumberMaps[8] = numberMaps[0]
		numberMaps = newNumberMaps
	}

	result := 0

	for _, count := range numberMaps {
		result += count
	}

	return result
}

func partTwo(scanner bufio.Scanner) int {

	var numberMaps = make(map[int]int)

	for scanner.Scan() {
		lineText := scanner.Text()
		numberStringTokens := strings.Split(lineText, ",")
		for _, numberString := range numberStringTokens {
			number, _ := strconv.Atoi(numberString)
			numberMaps[number]++
		}
	}

	for i := 0; i < 256; i++ {
		newNumberMaps := make(map[int]int)
		newNumberMaps[0] = numberMaps[1]
		newNumberMaps[1] = numberMaps[2]
		newNumberMaps[2] = numberMaps[3]
		newNumberMaps[3] = numberMaps[4]
		newNumberMaps[4] = numberMaps[5]
		newNumberMaps[5] = numberMaps[6]
		newNumberMaps[6] = numberMaps[7] + numberMaps[0]
		newNumberMaps[7] = numberMaps[8]
		newNumberMaps[8] = numberMaps[0]
		numberMaps = newNumberMaps
	}

	result := 0

	for _, count := range numberMaps {
		result += count
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
	result := partOne(*scanner)

	fmt.Println(result)
}
