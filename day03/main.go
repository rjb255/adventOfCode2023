package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())

}

func part1() int {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	result := 0
	var contents [][]byte
	re := regexp.MustCompile("[0-9]+")
	for scanner.Scan() {
		line := scanner.Text()
		contents = append(contents, []byte(line))
	}
	for lineNumber, line := range contents {
		indexes := re.FindAllIndex(line, -1)
		for _, index := range indexes {
			result += value(contents, lineNumber, index)
		}
	}
	return result
}

func value(content [][]byte, lineNumber int, numberRange []int) int {
	minLine := max(lineNumber-1, 0)
	maxLine := min(lineNumber+1, len(content)-1)
	minInput := max(numberRange[0]-1, 0)
	maxInput := min(numberRange[1]+1, len(content[0]))
	value, err := strconv.Atoi(string(content[lineNumber][numberRange[0]:numberRange[1]]))
	check(err)

	lines := content[minLine : maxLine+1]

	rejects := regexp.MustCompile("[^0-9\\.]")
	returnValue := 0
	for i := range lines {
		a := lines[i][minInput:maxInput]
		if rejects.Match(a) {
			returnValue = value
		}
	}
	return returnValue
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func max(num1 int, num2 int) int {
	return -1 * min(-1*num1, -1*num2)
}

func part2() int {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	result := 0
	var contents [][]byte
	re := regexp.MustCompile("\\*")
	for scanner.Scan() {
		line := scanner.Text()
		contents = append(contents, pad([]byte(line)))
	}
	contents = append([][]byte{bottom(len(contents[0]))}, contents...)
	contents = append(contents, bottom(len(contents[0])))

	for lineNumber, line := range contents {
		indexes := re.FindAllIndex(line, -1)
		for _, index := range indexes {
			result += gearValue(contents, lineNumber, index[0])
		}
	}
	return result
}

func gearValue(content [][]byte, lineNumber int, index int) int {
	product := 1
	gearCount := 0
	if unicode.IsDigit(rune(content[lineNumber-1][index])) {
		gearCount++
		product *= getNumber(content, lineNumber-1, index)
	} else {
		if unicode.IsDigit(rune(content[lineNumber-1][index-1])) {
			gearCount++
			product *= getNumber(content, lineNumber-1, index-1)
		}
		if unicode.IsDigit(rune(content[lineNumber-1][index+1])) {
			gearCount++
			product *= getNumber(content, lineNumber-1, index+1)
		}
	}
	if unicode.IsDigit(rune(content[lineNumber][index-1])) {
		gearCount++
		product *= getNumber(content, lineNumber, index-1)
	}
	if unicode.IsDigit(rune(content[lineNumber][index+1])) {
		gearCount++
		product *= getNumber(content, lineNumber, index+1)
	}
	if unicode.IsDigit(rune(content[lineNumber+1][index])) {
		gearCount++
		product *= getNumber(content, lineNumber+1, index)
	} else {
		if unicode.IsDigit(rune(content[lineNumber+1][index-1])) {
			product *= getNumber(content, lineNumber+1, index-1)
			gearCount++
		}
		if unicode.IsDigit(rune(content[lineNumber+1][index+1])) {
			gearCount++
			product *= getNumber(content, lineNumber+1, index+1)
		}
	}
	if gearCount == 2 {
		return product
	}

	return 0
}

func left(content [][]byte, lineNumber int, index int) string {
	if unicode.IsDigit(rune(content[lineNumber][index])) {
		return left(content, lineNumber, index-1) + string(content[lineNumber][index])
	}
	return ""
}

func right(content [][]byte, lineNumber int, index int) string {
	if unicode.IsDigit(rune(content[lineNumber][index])) {
		return string(content[lineNumber][index]) + right(content, lineNumber, index+1)
	}
	return ""
}

func getNumber(content [][]byte, lineNumber int, index int) int {
	number, err := strconv.Atoi(left(content, lineNumber, index-1) + string(content[lineNumber][index]) + right(content, lineNumber, index+1))
	check(err)
	return number
}

func pad(line []byte) []byte {
	return append([]byte("."), append(line, []byte(".")...)...)
}

func bottom(length int) []byte {
	var line []byte
	for i := 0; i < length; i++ {
		line = append(line, []byte(".")...)
	}
	return line
}
