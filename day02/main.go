package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/exp/maps"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1()
	part2()

}

func part1() {
	balls := make(map[string]int)
	balls["red"] = 12
	balls["green"] = 13
	balls["blue"] = 14

	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("((?P<number>[0-9]+)|(?P<colour>[a-z]+))")
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums_and_colours := re.FindAllString(line, -1)
		game_number, err := strconv.Atoi(nums_and_colours[1])
		check(err)
		total += game_number
		for i := 2; i < len(nums_and_colours); i += 2 {
			num, err := strconv.Atoi(nums_and_colours[i])
			check(err)
			if balls[nums_and_colours[i+1]] < num {
				total -= game_number
				break
			}
		}

	}
	fmt.Println(total)
}

func part2() {

	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("((?P<number>[0-9]+)|(?P<colour>[a-z]+))")
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums_and_colours := re.FindAllString(line, -1)
		balls := make(map[string]int)
		for i := 2; i < len(nums_and_colours); i += 2 {
			num, err := strconv.Atoi(nums_and_colours[i])
			check(err)
			if balls[nums_and_colours[i+1]] < num {
				balls[nums_and_colours[i+1]] = num
			}
		}
		values := maps.Values(balls)
		if len(values) > 0 {
			subtotal := 1
			for _, value := range values {
				subtotal *= value
			}
			total += subtotal
		}

	}
	fmt.Println(total)
}
