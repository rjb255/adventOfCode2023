package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1("test.txt")
	part2("test.txt")

	part1("input.txt")
	part2("input.txt")
}

func part1(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	numbersRegex := regexp.MustCompile("-{0,1}\\d+")
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbersText := numbersRegex.FindAllString(line, -1)

		numbers := []int{}

		for _, number := range numbersText {
			number, err := strconv.Atoi(number)
			check(err)
			numbers = append(numbers, number)
			numbers = append([]int{number}, numbers...)
		}
		numbers = nextInSequence(numbers)
		sum += numbers[len(numbers)-1]
	}
	println(sum)
}

func part2(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	numbersRegex := regexp.MustCompile("-{0,1}\\d+")
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbersText := numbersRegex.FindAllString(line, -1)

		numbers := []int{}

		for _, number := range numbersText {
			number, err := strconv.Atoi(number)
			check(err)
			numbers = append([]int{number}, numbers...)
		}
		numbers = nextInSequence(numbers)
		sum += numbers[len(numbers)-1]
	}
	println(sum)
}

func nextInSequence(numbers []int) []int {
	allSame := true
	for _, s := range numbers {
		if s != numbers[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return append(numbers, numbers[0])
	}
	diff := []int{}
	for i := 1; i < len(numbers); i++ {
		diff = append(diff, numbers[i]-numbers[i-1])
	}
	diff = nextInSequence(diff)
	return append(numbers, numbers[len(numbers)-1]+diff[len(diff)-1])
}
