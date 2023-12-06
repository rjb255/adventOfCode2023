package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var numberRegex = regexp.MustCompile("\\d+")

func main() {
	file, err := os.Open("day5.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	seeds := []int{}
	maps := [7][][]int{}
	modifiedSeeds := []int{}
	mapCounter := -1
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if lineNumber == 1 {
			seeds = lineToNumbers(line)
			for i := 0; i < len(seeds); i += 2 {
				for j := 0; j < seeds[i+1]; j++ {
					modifiedSeeds = append(modifiedSeeds, seeds[i]+j)
				}
			}
			continue
		}
		if regexp.MustCompile("[a-zA-Z]").MatchString(line) {
			mapCounter++
			continue
		}
		if numberRegex.MatchString(line) {
			pointers := [2]int{0, len(maps[mapCounter])}
			maps[mapCounter] = insertLine(maps[mapCounter], lineToNumbers(line), pointers)
		}
	}
	minLocation := 346433842
	for _, seed := range seeds {
		location := seed2area(maps[:], seed)
		if location < minLocation {
			minLocation = location
		}
	}
	fmt.Println(minLocation)
	for i := 0; i < len(seeds); i += 2 {
		fmt.Printf("SeedCat: %d\n", seeds[i])
		for j := 0; j < seeds[i+1]; j++ {
			location := seed2area(maps[:], seeds[i]+j)
			if location < minLocation {
				minLocation = location
				fmt.Printf("New Min: %d\n", location)
			}
		}
	}

	fmt.Println(minLocation)
}

func lineToNumbers(line string) []int {
	matches := numberRegex.FindAllString(line, -1)
	var values []int
	for _, match := range matches {
		value, err := strconv.Atoi(match)
		check(err)
		values = append(values, value)
	}
	return values
}

func seed2area(maps [][][]int, value int) int {
	if len(maps) == 0 {
		return value
	}
	currentMap := maps[0]
	for _, entry := range currentMap {
		// fmt.Println(entry)
		if entry[1] <= value && value <= entry[1]+entry[2] {
			newValue := entry[0] + (value - entry[1])
			// fmt.Printf("Value %d New %d\n", value, newValue)
			return seed2area(maps[1:], newValue)
		}
	}
	// fmt.Printf("Value %d New %d\n", value, value)
	return seed2area(maps[1:], value)
}

func insertLine(mapping [][]int, line []int, pointers [2]int) [][]int {
	return append(mapping, line)
	// mid := pointers[0] + (pointers[1]-pointers[0])/2
	// if pointers[0] >= pointers[1] {
	// 	mapping = append(append(mapping[:mid], line), mapping[mid:]...)
	// 	return mapping
	// }
	// if mapping[mid][1] < line[1] {
	// 	pointers[0] = mid + 1
	// } else {
	// 	pointers[1] = mid - 1
	// }
	// return insertLine(mapping, line, pointers)
}
