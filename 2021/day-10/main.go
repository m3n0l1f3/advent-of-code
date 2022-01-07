package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
)

var SyntaxErrorScoreMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var AutocompleteScoreMap = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func validateLine(line string) (valid bool, corrupt rune, remaining []rune) {

	l := list.New()

	for _, r := range line {
		lastElement := l.Back()
		// Empty queue
		if lastElement == nil {
			l.PushBack(r)
		} else if r == '{' || r == '[' || r == '<' || r == '(' {
			l.PushBack(r)
		} else if (lastElement.Value == '{' && r == '}') || (lastElement.Value == '<' && r == '>') || (lastElement.Value == '[' && r == ']') || (lastElement.Value == '(' && r == ')') {
			l.Remove(lastElement)
		} else {
			valid = false
			corrupt = r
			return
		}
	}

	valid = true
	remaining = make([]rune, 0)
	for e := l.Back(); e != nil; e = e.Prev() {
		r := e.Value.(rune)
		remaining = append(remaining, r)
	}

	return
}

func partOne(scanner bufio.Scanner) int {

	result := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		valid, corrupt, _ := validateLine(lineText)
		if !valid {
			result += SyntaxErrorScoreMap[corrupt]
		}
	}

	return result
}

func partTwo(scanner bufio.Scanner) int {

	scoreRank := make([]int, 0)
	for scanner.Scan() {
		lineText := scanner.Text()
		valid, _, remaining := validateLine(lineText)
		if valid {
			lineScore := 0
			for _, r := range remaining {
				lineScore *= 5
				lineScore += AutocompleteScoreMap[r]
			}

			scoreRank = append(scoreRank, lineScore)
		}
	}

	sort.Ints(scoreRank)
	return scoreRank[(len(scoreRank))/2]

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
