package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("PART 1 Test:")
	part1("test.txt")
	fmt.Println("PART 2 Test:")
	part2("test.txt")

	fmt.Println("PART 1:")
	part1("input.txt")
	fmt.Println("PART 2:")
	part2("input.txt")
}

func part1(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	galaxyRegex := regexp.MustCompile("#")
	lineNumber := -1
	xs := []int{}
	ys := []int{}
	totalGalaxyCount := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		galaxies := galaxyRegex.FindAllStringIndex(line, -1)
		if len(galaxies) == 0 {
			lineNumber++
			continue
		}
		for _, galaxy := range galaxies {
			ys = append(ys, lineNumber)
			xs = append(xs, galaxy[0])
			totalGalaxyCount++
		}
	}

	tot := 0
	for i, y := range ys {
		tot -= (totalGalaxyCount - 1 - 2*i) * y
	}
	slices.Sort(xs)
	offset := 0
	println(tot)
	for i, x := range xs {
		if i > 0 && xs[i]-xs[i-1] > 1 {
			offset += xs[i] - xs[i-1] - 1
		}
		tot -= (totalGalaxyCount - 1 - 2*i) * (x + offset)
	}
	println(tot)
}
func part2(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	galaxyRegex := regexp.MustCompile("#")
	lineNumber := -1
	xs := []int{}
	ys := []int{}
	totalGalaxyCount := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		galaxies := galaxyRegex.FindAllStringIndex(line, -1)
		if len(galaxies) == 0 {
			lineNumber += 1000000 - 1
			continue
		}
		for _, galaxy := range galaxies {
			ys = append(ys, lineNumber)
			xs = append(xs, galaxy[0])
			totalGalaxyCount++
		}
	}

	tot := 0
	for i, y := range ys {
		tot -= (totalGalaxyCount - 1 - 2*i) * y
	}

	slices.Sort(xs)
	offset := 0
	for i, x := range xs {
		if i > 0 && xs[i]-xs[i-1] > 1 {
			offset += (1000000 - 1) * (xs[i] - xs[i-1] - 1)
		}
		tot -= (totalGalaxyCount - 1 - 2*i) * (x + offset)
	}
	println(tot)
}
