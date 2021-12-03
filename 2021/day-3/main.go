package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne(scanner bufio.Scanner, bitLength int) int {

	bitMap := make([]int, bitLength)

	for scanner.Scan() {
		binaryText := scanner.Text()
		for index, bit := range binaryText {
			if bit == '0' {
				bitMap[index]--
			} else {
				bitMap[index]++
			}
		}
	}

	var gammaRateBuilder strings.Builder
	var epsilonRateBuilder strings.Builder

	for index, count := range bitMap {
		if count > 0 {
			gammaRateBuilder.WriteRune('1')
			epsilonRateBuilder.WriteRune('0')
		} else if count < 0 {
			gammaRateBuilder.WriteRune('0')
			epsilonRateBuilder.WriteRune('1')
		} else {
			fmt.Println("Help! :", index)
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateBuilder.String(), 2, 32)
	epsilonRate, _ := strconv.ParseInt(epsilonRateBuilder.String(), 2, 32)

	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
		os.Exit(1)
	}

	return int(gammaRate * epsilonRate)
}

func findOxygenGeneratorRating(bitStringList []string, bitLength int) int {
	localStringList := bitStringList

	for i := 0; i < bitLength; i++ {
		count := 0
		newLocalStringList := make([]string, 0)
		for _, localString := range localStringList {
			if localString[i] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		for _, localString := range localStringList {
			if count >= 0 && localString[i] == '1' {
				newLocalStringList = append(newLocalStringList, localString)
			} else if count < 0 && localString[i] == '0' {
				newLocalStringList = append(newLocalStringList, localString)
			}
		}

		localStringList = newLocalStringList
		if len(localStringList) == 1 {
			break
		}
	}

	result, _ := strconv.ParseInt(localStringList[0], 2, 64)
	return int(result)
}

func findCO2GeneratorRating(bitStringList []string, bitLength int) int {
	localStringList := bitStringList

	for i := 0; i < bitLength; i++ {
		count := 0
		newLocalStringList := make([]string, 0)
		for _, localString := range localStringList {
			if localString[i] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		for _, localString := range localStringList {
			if count >= 0 && localString[i] == '0' {
				newLocalStringList = append(newLocalStringList, localString)
			} else if count < 0 && localString[i] == '1' {
				newLocalStringList = append(newLocalStringList, localString)
			}
		}

		localStringList = newLocalStringList
		if len(localStringList) == 1 {
			break
		}
	}

	result, _ := strconv.ParseInt(localStringList[0], 2, 64)
	return int(result)
}

func partTwo(scanner bufio.Scanner, bitLength int) int {

	bitStringList := []string{}
	for scanner.Scan() {
		binaryText := scanner.Text()
		bitStringList = append(bitStringList, binaryText)
	}

	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
		os.Exit(1)
	}

	oxygenRating := findOxygenGeneratorRating(bitStringList, bitLength)
	co2Rating := findCO2GeneratorRating(bitStringList, bitLength)

	return int(oxygenRating * co2Rating)
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

	result := partTwo(*scanner, 12)

	fmt.Println(result)
}
