package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

var previousReached = 1

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part2("day8.test.txt")

	part1("day8.txt")
	part2("day8.txt")
}

func part1(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	lineNumber := -1
	instructions := []string{}
	content := make(map[string][]string)

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if lineNumber == 0 {
			for _, content := range line {
				instructions = append(instructions, string(content))
			}
			continue
		}
		if lineNumber > 1 {
			text := regexp.MustCompile("[A-Z]{3}").FindAllString(line, -1)
			content[text[0]] = text[1:]
		}
	}
	nextInstruction := "AAA"
	i := 0
	for nextInstruction != "ZZZ" {
		if instructions[i%len(instructions)] == "L" {
			nextInstruction = content[nextInstruction][0]
		} else {
			nextInstruction = content[nextInstruction][1]
		}
		i++
	}
	fmt.Println(i)
}

func checkSlice(slice []string) bool {
	for i, loc := range slice {
		if string(loc[2]) != "Z" {
			if i >= previousReached {
				previousReached = i
				fmt.Println("Reached: ", i)
			}
			return false
		}
	}
	return true
}

func part2(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	lineNumber := -1
	instructions := []string{}
	content := make(map[string][]string)
	nextInstructions := []string{}

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if lineNumber == 0 {
			for _, content := range line {
				instructions = append(instructions, string(content))
			}
			continue
		}
		if lineNumber > 1 {
			text := regexp.MustCompile("[0-9A-Z]{3}").FindAllString(line, -1)
			content[text[0]] = text[1:]
			if string(text[0][2]) == "A" {
				nextInstructions = append(nextInstructions, text[0])
			}
		}
	}

	cycles := [][]int{}
	counts := [][]uint64{}

	for _, nextInstruction := range nextInstructions {
		history := []string{nextInstruction}
		for {
			cycle, offset, count := checkHistory(history, len(instructions))
			if cycle > 0 {
				cycles = append(cycles, []int{cycle, offset})
				for i := range count {
					count[i] += uint64(offset)
				}
				counts = append(counts, count)
				break
			}
			if instructions[(len(history)-1)%len(instructions)] == "L" {
				nextInstruction = content[nextInstruction][0]
			} else {
				nextInstruction = content[nextInstruction][1]
			}
			history = append([]string{nextInstruction}, history...)
		}
	}

	for !match(counts) {
		for i := 0; i < len(counts); i++ {

			m1, i1 := minSlice(counts[i])
			m2, _ := minSlice(counts[(i+1)%len(counts)])
			if m1 < m2 {

				counts[i][i1] += uint64(cycles[i][0])
			}
		}
	}

	product := 1.
	for _, s := range cycles {
		product *= float64(s[0])
	}

	fmt.Println(slices.Min(counts[0]))
}

func checkHistory(history []string, rollover int) (int, int, []uint64) {
	if len(history) < rollover {
		return 0, 0, []uint64{}
	}
	for i := rollover; i < len(history); i += rollover {
		if history[0] == history[i] {
			zs := []uint64{}
			for j := 0; j < i; j++ {
				if string(history[j][2]) == "Z" {
					zs = append(zs, uint64(i-j))
				}
			}
			fmt.Println(i, (len(history)-1)%rollover, zs)
			return i, (len(history) - 1) % rollover, zs
		}
	}
	return 0, 0, []uint64{}
}

func match(counts [][]uint64) bool {
	for _, count := range counts[0] {
		_match := true
		for i := 1; i < len(counts); i++ {
			if !slices.Contains(counts[i], count) {
				_match = false
				break
			}
		}
		if _match {
			return true
		}
	}
	return false
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func minUnit64(n1, n2 uint64) uint64 {
	if n1 < n2 {
		return n1
	}
	return n2
}

func minSlice(s []uint64) (uint64, int) {
	m := s[0]
	i := 0

	for index, v := range s {
		if _min := minUnit64(m, v); _min != m {
			i = index
			m = _min
		}
	}
	return m, i
}
