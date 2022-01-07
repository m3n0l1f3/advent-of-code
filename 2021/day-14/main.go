package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func loadData(scanner bufio.Scanner) (string, map[string]string) {

	var initialString string
	templateMap := make(map[string]string)

	for scanner.Scan() {
		lineText := scanner.Text()
		if len(lineText) == 0 {
			// no-op
		} else if strings.Contains(lineText, "->") {
			token := strings.Split(lineText, " -> ")
			templateMap[token[0]] = token[1]
		} else {
			initialString = lineText
		}
	}

	return initialString, templateMap
}

func mutate(initialString string, templateMap map[string]string) string {

	totalLength := len(initialString)
	newString := ""
	for i := 0; i < totalLength-1; i++ {
		tokenString := initialString[i : i+2]
		token := templateMap[tokenString]
		newString = newString + string(initialString[i]) + token
	}

	newString = newString + string(initialString[totalLength-1])

	return newString
}

func intMax(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

func intMin(left int, right int) int {
	if left < right {
		return left
	}
	return right
}

func mutateTwo(pairMap map[string]int, templateMap map[string]string) map[string]int {

	newPairMap := make(map[string]int)
	for pair, count := range pairMap {
		token, ok := templateMap[pair]
		if ok {
			tokenOne := string(pair[0]) + token
			tokenTwo := token + string(pair[1])
			newPairMap[tokenOne] += count
			newPairMap[tokenTwo] += count
		} else {
			newPairMap[pair] += count
		}
	}

	return newPairMap
}

func partOne(initialString string, templateMap map[string]string, steps int) int {

	for i := 0; i < steps; i++ {
		initialString = mutate(initialString, templateMap)
	}

	countMap := make(map[rune]int)
	for _, r := range initialString {
		countMap[r]++
	}

	minCount, maxCount := math.MaxInt, math.MinInt

	for _, value := range countMap {
		minCount = intMin(minCount, value)
		maxCount = intMax(maxCount, value)
	}

	return maxCount - minCount
}

func partTwo(initialString string, templateMap map[string]string, steps int) int {

	pairMap := make(map[string]int)
	totalLength := len(initialString)
	for i := 0; i < totalLength-1; i++ {
		pairMap[initialString[i:i+2]]++
	}

	for i := 0; i < steps; i++ {
		newPairMap := mutateTwo(pairMap, templateMap)
		pairMap = newPairMap
	}

	countMap := make(map[rune]int)
	for token, count := range pairMap {
		countMap[rune(token[0])] += count
		countMap[rune(token[1])] += count
	}

	minCount, maxCount := math.MaxInt, math.MinInt

	for _, value := range countMap {
		minCount = intMin(minCount, value)
		maxCount = intMax(maxCount, value)
	}

	return (maxCount - minCount + 1) / 2
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
	initialString, templateMap := loadData(*scanner)
	result := partTwo(initialString, templateMap, 40)

	fmt.Println(result)
}
