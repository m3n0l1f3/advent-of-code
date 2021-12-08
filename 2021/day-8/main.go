package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func partOne(scanner bufio.Scanner) int {

	count := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		inputTokens := strings.Split(lineText, " | ")
		outputString := inputTokens[1]
		outputTokens := strings.Split(outputString, " ")

		for _, token := range outputTokens {
			// `1` --> 2
			// `4` --> 4
			// `7` --> 3
			// `8` --> 7

			switch len(token) {
			case 2:
				count++
			case 4:
				count++
			case 3:
				count++
			case 7:
				count++
			}
		}
	}

	return count
}

// Sort By String Length
type byLen []string

func (a byLen) Len() int {
	return len(a)
}

func (a byLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a byLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func diffStrings(a string, b string) (aDiff string, bDiff string) {
	aDiff = ""
	for _, aRune := range a {
		if !strings.ContainsRune(b, aRune) {
			aDiff += string(aRune)
		}
	}

	bDiff = ""
	for _, bRune := range b {
		if !strings.ContainsRune(a, bRune) {
			bDiff += string(bRune)
		}
	}

	return
}

// Only 0 doesn't have d, compare amongst all 5 & 6 to find 0
func findDigitZero(targetA string, targetBD string, slices []string) int {
	firstThree := slices[0:3]
	countMap := make(map[string]int)
	for _, token := range firstThree {
		for _, r := range token {
			countMap[string(r)]++
		}
	}

	targetD := ""
	for key, value := range countMap {
		if key == targetA {
			continue
		}
		if value == 3 && strings.Contains(targetBD, key) {
			targetD = key
		}
	}

	for index, token := range slices[3:] {
		if !strings.Contains(token, targetD) {
			return index
		}
	}
	return 0
}

func findDigitSix(digitOne string, slices []string) int {
	for index, token := range slices {
		for _, digit := range digitOne {
			if !strings.ContainsRune(token, digit) {
				return index
			}
		}
	}
	return 0
}

func findDigitNine(slices []int) int {
	for _, i := range []int{6, 7, 8} {
		found := false
		for _, number := range slices {
			if i == number {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return 0
}

func findDigitFive(targetB string, slices []string) int {
	for index, token := range slices {
		if strings.Contains(token, targetB) {
			return index
		}
	}
	return 0
}

func findDigitTwo(targetE string, slices []string) int {
	for index, token := range slices {
		if strings.Contains(token, targetE) {
			return index
		}
	}
	return 0
}

func findDigitThree(slices []int) int {
	for _, i := range []int{3, 4, 5} {
		found := false
		for _, number := range slices {
			if i == number {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return 0
}

func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

func partTwo(scanner bufio.Scanner) int {

	count := 0
	for scanner.Scan() {
		digitMap := make(map[string]int)
		lineText := scanner.Text()
		inputTokens := strings.Split(lineText, " | ")
		mapString := inputTokens[0]
		mapTokens := strings.Split(mapString, " ")
		sort.Sort(byLen(mapTokens))

		// 0 --> a b c   e f g   6
		// 1 -->     c     f     2
		// 2 --> a   c d e   g   5
		// 3 --> a   c d   f g   5
		// 4 -->   b c d   f     4
		// 5 --> a b   d   f g   5
		// 6 --> a b   d e f g   6
		// 7 --> a   c     f     3
		// 8 --> a b c d e f g   7
		// 9 --> a b c d   f g   6

		// 0 1 2 3 ~ 5 6 ~ 8 9
		// 1 7 4 2,3,5 0,6,9 8
		digitMap[mapTokens[0]] = 1
		digitMap[mapTokens[1]] = 7
		digitMap[mapTokens[2]] = 4
		digitMap[mapTokens[9]] = 8

		_, targetA := diffStrings(mapTokens[0], mapTokens[1])
		targetBD, _ := diffStrings(mapTokens[2], mapTokens[1])

		indexZero := findDigitZero(targetA, targetBD, mapTokens[3:9]) + 6
		digitMap[mapTokens[indexZero]] = 0

		_, targetD := diffStrings(mapTokens[indexZero], mapTokens[9])
		targetB, _ := diffStrings(targetBD, targetD)

		indexSix := findDigitSix(mapTokens[0], mapTokens[6:9]) + 6
		digitMap[mapTokens[indexSix]] = 6

		indexNine := findDigitNine([]int{indexZero, indexSix})
		digitMap[mapTokens[indexNine]] = 9

		indexFive := findDigitFive(targetB, mapTokens[3:6]) + 3
		digitMap[mapTokens[indexFive]] = 5

		targetE, _ := diffStrings(mapTokens[indexZero], mapTokens[indexNine])
		indexTwo := findDigitTwo(targetE, mapTokens[3:6]) + 3
		digitMap[mapTokens[indexTwo]] = 2

		indexThree := findDigitThree([]int{indexTwo, indexFive})
		digitMap[mapTokens[indexThree]] = 3

		sortedMap := make(map[string]int)
		for key, value := range digitMap {
			sortedMap[sortString(key)] = value
		}

		outputString := inputTokens[1]
		outputTokens := strings.Split(outputString, " ")
		outputDigit := 0
		for _, token := range outputTokens {
			token = sortString(token)
			digit := sortedMap[token]
			outputDigit = outputDigit*10 + digit
		}
		count += outputDigit
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
	result := partTwo(*scanner)

	fmt.Println(result)
}
