package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var cardMapping = make(map[string]int)

var gaps = [7]float64{
	math.Pow(14, 7),
	math.Pow(14, 7),
	math.Pow(14, 7),
	math.Pow(14, 7),
	math.Pow(14, 7),
	math.Pow(14, 7),
	math.Pow(14, 7),
}

var boundaries = [7]float64{}

func fillBoundaries(index int) float64 {
	if index <= 0 {
		return gaps[0]
	}
	return gaps[index] + fillBoundaries(index-1)
}

func main() {
	for i := 0; i < len(gaps); i++ {
		boundaries[i] = fillBoundaries(i)
	}
	for i := 2; i < 10; i++ {
		cardMapping[strconv.Itoa(i)] = i
	}
	cardMapping["T"] = 10
	cardMapping["J"] = 11
	cardMapping["Q"] = 12
	cardMapping["K"] = 13
	cardMapping["A"] = 14

	part1("day7.test.txt")
	part1("day7.txt")
}

func part1(fileName string) {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)

	sortedUseful := [][]float64{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := (strings.Fields(line))
		score, err := strconv.Atoi(fields[1])
		check(err)
		useful := []float64{cards2value(fields[0]), float64(score)}
		sortedUseful = insertLine(sortedUseful, useful)
	}
	score := 0.
	for i, row := range sortedUseful {
		score += float64(i+1) * row[1]
	}
	fmt.Println(int(score))

}

func cards2value(hand string) float64 {
	cards := []int{}
	for _, card := range hand {
		cards = append(cards, cardMapping[string(card)])
	}
	slices.Sort(cards)

	counts := []int{1}
	for index, card := range cards {
		if index == 0 {
			continue
		}
		if card == cards[index-1] {
			counts[len(counts)-1]++
		} else {
			counts = append(counts, 1)
		}
	}

	score := 0.

	j := 0

	counts, cards = sortCounts(counts, cards, 0, 0)
	for i := 0; i < len(counts); i++ {
		score += float64(cards[j]) * math.Pow(14, float64(i))
		j += counts[i]
	}

	sortedCounts := []int(counts)
	slices.Sort(sortedCounts)
	// High card, where all cards' labels are distinct: 23456
	// 14*13*12*11 = 240240
	if checkSame([]int{1, 1, 1, 1, 1}, sortedCounts) {
		return score
	}

	// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	// 14*13*12*11 = 24024
	if checkSame([]int{1, 1, 1, 2}, sortedCounts) {
		return boundaries[0] + score
	}

	// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	// 14*13*12 = 2184
	if checkSame([]int{1, 2, 2}, sortedCounts) {
		return boundaries[1] + score
	}

	// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	// 14*13*12 = 2184
	if checkSame([]int{1, 1, 3}, sortedCounts) {
		return boundaries[2] + score
	}

	// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	// 14*13 = 182
	if checkSame([]int{2, 3}, sortedCounts) {
		return boundaries[3] + score
	}
	// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	// 14*13 = 182
	if checkSame([]int{1, 4}, sortedCounts) {
		return boundaries[4] + score
	}
	// Five of a kind, where all five cards have the same label: AAAAA
	// 14 = 14
	return boundaries[5] + score
}

func checkSame(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func insertLine(scores [][]float64, line []float64) [][]float64 {
	pointer1 := 0
	pointer2 := len(scores)

	for {
		mid := pointer1 + (pointer2-pointer1)/2
		if pointer1 >= pointer2 {
			mid = pointer1
			if mid == len(scores) {
				return append(scores, line)
			}
			if mid == 0 {
				return append([][]float64{line}, scores...)
			}
			newScores := append([][]float64{}, scores[:mid]...)
			newScores = append(newScores, line)
			newScores = append(newScores, scores[mid:]...)
			return newScores
		}
		if scores[mid][0] > line[0] {
			pointer2 = max(mid, pointer1)
		} else {
			pointer1 = min(mid+1, pointer2)
		}
	}
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func max(num1, num2 int) int {
	return -min(-num1, -num2)
}

func sortCounts(counts, cards []int, countPoint, cardPoint int) ([]int, []int) {
	if countPoint == 0 {
		cardPoint += counts[countPoint]
		countPoint++
		return sortCounts(counts, cards, countPoint, cardPoint)
	}
	if countPoint >= len(counts) {
		return counts, cards
	}
	if counts[countPoint] < counts[countPoint-1] {
		start := cards[:cardPoint-counts[countPoint-1]]
		temp1 := cards[cardPoint-counts[countPoint-1] : cardPoint]
		chunk := cards[cardPoint : cardPoint+counts[countPoint]]
		end := cards[cardPoint+counts[countPoint]:]
		newCards := append(append(append(append([]int{}, start...), chunk...), temp1...), end...)
		cardPoint -= counts[countPoint-1]

		counts[countPoint], counts[countPoint-1] = counts[countPoint-1], counts[countPoint]
		countPoint--
		return sortCounts(counts, newCards, countPoint, cardPoint)
	}
	cardPoint += counts[countPoint]
	countPoint++
	return sortCounts(counts, cards, countPoint, cardPoint)
}
