package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne(scanner bufio.Scanner) int {

	var numberList = make([]int, 0)

	for scanner.Scan() {
		lineText := scanner.Text()
		numberStringTokens := strings.Split(lineText, ",")
		for _, numberString := range numberStringTokens {
			number, _ := strconv.Atoi(numberString)
			numberList = append(numberList, number)
		}
	}

	sort.Ints(numberList)
	listLength := len(numberList)
	medium := 0
	if listLength%2 == 0 {
		medium = (numberList[listLength/2-1] + numberList[listLength/2]) / 2
	} else {
		medium = numberList[listLength/2-1]
	}
	result := 0

	for _, number := range numberList {
		result += int(math.Abs(float64(number - medium)))
	}

	return result
}

func partTwo(scanner bufio.Scanner) int {

	var numberList = make([]int, 0)
	var totalSum int64 = 0

	for scanner.Scan() {
		lineText := scanner.Text()
		numberStringTokens := strings.Split(lineText, ",")
		for _, numberString := range numberStringTokens {
			number, _ := strconv.Atoi(numberString)
			numberList = append(numberList, number)
			totalSum += int64(number)
		}
	}

	average := int(float64(totalSum) / float64(len(numberList)))
	resultOne := 0
	for _, number := range numberList {
		diff := int(math.Abs(float64(number - average)))
		cost := int(float64(1+diff) / float64(2) * float64(diff))
		resultOne += cost
	}

	resultTwo := 0
	for _, number := range numberList {
		diff := int(math.Abs(float64(number - average + 1)))
		cost := int(float64(1+diff) / float64(2) * float64(diff))
		resultTwo += cost
	}

	return int(math.Min(float64(resultOne), float64(resultTwo)))
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
