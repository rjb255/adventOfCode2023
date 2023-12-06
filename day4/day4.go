package main

import (
	"bufio"
	"fmt"
	"math"
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
	file, err := os.Open("day4.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	groups := regexp.MustCompile("(:|\\|)")
	numbers := regexp.MustCompile("\\d+")
	total := 0
	var scoreTracker []int
	for scanner.Scan() {
		var winnings []int
		var actual []int

		line := scanner.Text()
		parts := groups.Split(line, -1)
		winningsText := numbers.FindAllString(parts[1], -1)
		actualText := numbers.FindAllString(parts[2], -1)
		for _, value := range winningsText {
			number, err := strconv.Atoi(value)
			winnings = append(winnings, number)
			check(err)
		}
		for _, value := range actualText {
			number, err := strconv.Atoi(value)
			actual = append(actual, number)
			check(err)
		}
		bubbleSort(winnings, 0)
		bubbleSort(actual, 0)
		score := winningsCount(winnings, actual)
		scoreTracker = append(scoreTracker, score)
		if score > 0 {
			total += int(math.Pow(2, float64(score-1)))
		}
	}
	fmt.Println(total)
	fmt.Println(part2Score(scoreTracker, true))

}

func bubbleSort(numbers []int, pointer int) {
	if len(numbers)-1 <= pointer {
		return
	}
	if numbers[pointer] > numbers[pointer+1] {
		temp := numbers[pointer]
		numbers[pointer] = numbers[pointer+1]
		numbers[pointer+1] = temp
		if pointer != 0 {
			bubbleSort(numbers, pointer-1)
			return
		}
	}
	bubbleSort(numbers, pointer+1)
	return
}

func winningsCount(winning []int, actual []int) int {
	pointer1 := 0
	pointer2 := 0
	count := 0
	for {
		if pointer1 >= len(winning) || pointer2 >= len(actual) {
			return count
		}
		if winning[pointer1] == actual[pointer2] {
			count++
			pointer2++
		} else if winning[pointer1] < actual[pointer2] {
			pointer1++
		} else {
			pointer2++
		}
	}
}

func part2Score(scores []int, main bool) int {
	if len(scores) == 0 {
		return 0
	}
	sum := 1

	for i := 1; i <= min(scores[0], len(scores)); i++ {
		sum += part2Score(scores[i:], false)
	}
	if main {
		sum += part2Score(scores[1:], true)
	}
	return sum
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
