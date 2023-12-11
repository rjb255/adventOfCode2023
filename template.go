package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("PART 1 Test:")
	part1("test.txt")
	fmt.Println("\nPART 2 Test:")
	// part2("test.txt")

	fmt.Println("\nPART 1:")
	part1("input.txt")
	fmt.Println("\nPART 2:")
	// part2("input.txt")
}

func part1(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}
	file.Close()
}
