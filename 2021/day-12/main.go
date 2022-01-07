package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cave struct {
	small      bool
	visitCount int
	neighbor   []string
}

func dedupePath(pathList [][]string) [][]string {

	dedupeMap := make(map[string]bool)
	result := make([][]string, 0)
	for _, path := range pathList {
		pathString := strings.Join(path, ",")
		if _, value := dedupeMap[pathString]; !value {
			dedupeMap[pathString] = true
			result = append(result, path)
		}
	}

	return result
}

func traverse(caveMap map[string]cave, path []string) [][]string {
	name := path[len(path)-1]
	if name == "end" {
		return [][]string{path}
	}
	caveInfo := caveMap[name]

	// Skip visited small cave
	if caveInfo.small && caveInfo.visitCount > 0 {
		return make([][]string, 0)
	}

	// Update count
	caveInfo.visitCount++
	caveMap[name] = caveInfo

	result := make([][]string, 0)
	for _, node := range caveInfo.neighbor {
		newPath := make([]string, len(path))
		copy(newPath, path)
		newPath = append(newPath, node)
		possiblePath := traverse(caveMap, newPath)

		result = append(result, possiblePath...)
	}

	caveInfo.visitCount--
	caveMap[name] = caveInfo

	return dedupePath(result)
}

func traverseWithAllowrance(caveMap map[string]cave, path []string, allowance int) [][]string {
	name := path[len(path)-1]
	if name == "end" {
		return [][]string{path}
	}
	caveInfo := caveMap[name]

	// Skip visited small cave
	if caveInfo.small {
		if caveInfo.visitCount > 0 {

			if name == "start" || allowance <= 0 {
				return make([][]string, 0)
			} else {
				allowance--
			}
		}
	}

	// Update count
	caveInfo.visitCount++
	caveMap[name] = caveInfo

	result := make([][]string, 0)
	for _, node := range caveInfo.neighbor {
		newPath := make([]string, len(path))
		copy(newPath, path)
		newPath = append(newPath, node)
		possiblePath := traverseWithAllowrance(caveMap, newPath, allowance)
		result = append(result, possiblePath...)
	}

	caveInfo.visitCount--
	caveMap[name] = caveInfo

	return dedupePath(result)
}

func partOne(scanner bufio.Scanner) int {

	caveMap := make(map[string]cave)
	for scanner.Scan() {
		lineText := scanner.Text()
		caveNames := strings.Split(lineText, "-")
		caveNameOne, caveNameTwo := caveNames[0], caveNames[1]

		if _, contained := caveMap[caveNameOne]; !contained {
			smallCave := strings.ToLower(caveNameOne) == caveNameOne
			caveMap[caveNameOne] = cave{smallCave, 0, make([]string, 0)}
		}

		if _, contained := caveMap[caveNameTwo]; !contained {
			smallCave := strings.ToLower(caveNameTwo) == caveNameTwo
			caveMap[caveNameTwo] = cave{smallCave, 0, make([]string, 0)}
		}

		caveOne := caveMap[caveNameOne]
		caveOne.neighbor = append(caveOne.neighbor, caveNameTwo)
		caveMap[caveNameOne] = caveOne

		caveTwo := caveMap[caveNameTwo]
		caveTwo.neighbor = append(caveTwo.neighbor, caveNameOne)
		caveMap[caveNameTwo] = caveTwo
	}

	path := []string{"start"}
	paths := traverse(caveMap, path)

	return len(paths)
}

func partTwo(scanner bufio.Scanner) int {

	caveMap := make(map[string]cave)
	for scanner.Scan() {
		lineText := scanner.Text()
		caveNames := strings.Split(lineText, "-")
		caveNameOne, caveNameTwo := caveNames[0], caveNames[1]

		if _, contained := caveMap[caveNameOne]; !contained {
			smallCave := strings.ToLower(caveNameOne) == caveNameOne
			caveMap[caveNameOne] = cave{smallCave, 0, make([]string, 0)}
		}

		if _, contained := caveMap[caveNameTwo]; !contained {
			smallCave := strings.ToLower(caveNameTwo) == caveNameTwo
			caveMap[caveNameTwo] = cave{smallCave, 0, make([]string, 0)}
		}

		caveOne := caveMap[caveNameOne]
		caveOne.neighbor = append(caveOne.neighbor, caveNameTwo)
		caveMap[caveNameOne] = caveOne

		caveTwo := caveMap[caveNameTwo]
		caveTwo.neighbor = append(caveTwo.neighbor, caveNameOne)
		caveMap[caveNameTwo] = caveTwo
	}

	path := []string{"start"}
	paths := traverseWithAllowrance(caveMap, path, 1)

	return len(paths)
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
