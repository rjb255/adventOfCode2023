package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var setMap = map[rune][][]int{
	'|': {{-2, -1, -2}, {-2, -1, -2}, {-2, -1, -2}},
	'-': {{-2, -2, -2}, {-1, -1, -1}, {-2, -2, -2}},
	'F': {{-2, -2, -2}, {-2, -1, -1}, {-2, -1, -2}},
	'7': {{-2, -2, -2}, {-1, -1, -2}, {-2, -1, -2}},
	'J': {{-2, -1, -2}, {-1, -1, -2}, {-2, -2, -2}},
	'L': {{-2, -1, -2}, {-2, -1, -1}, {-2, -2, -2}},
	'.': {{-2, -2, -2}, {-2, -1, -2}, {-2, -2, -2}},
	'S': {{-2, -1, -2}, {-1, 0, -1}, {-2, -1, -2}},
}

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
	content := readFullFile(filename)
	distances := expandContent(content)
	distances = padContent(distances)
	x, y := findStart(distances)
	explore(distances, x, y)
	// for _, line := range distances {
	// 	for _, char := range line {
	// 		fmt.Printf("%-8v", char)
	// 	}
	// 	fmt.Println("")
	// }
	fmt.Println(findLargest(distances) / 3)
}

func part2(filename string) {
	content := readFullFile(filename)
	distances := expandContent(content)
	distances = padContent(distances)
	x, y := findStart(distances)
	explore(distances, x, y)
	inverted := invert(distances)
	scan(inverted)
	// for _, line := range inverted {
	// 	for _, char := range line {
	// 		fmt.Printf("%-8v", char)
	// 	}
	// 	fmt.Println("")
	// }
	tot := 0
	for i := 2; i < len(inverted); i += 3 {
		for j := 2; j < len(inverted[0]); j += 3 {
			if inverted[i][j] > 0 {
				tot++
			}
		}
	}
	fmt.Println(tot)
}

func readFullFile(filename string) [][]rune {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	content := [][]rune{}
	for scanner.Scan() {
		content = append(content, []rune(scanner.Text()))
	}
	return content
}

func findStart(content [][]int) (x, y int) {
	for i := 0; i < len(content); i++ {
		for j := 0; j < len(content[0]); j++ {
			if content[i][j] == 0 {
				return i, j
			}
		}
	}
	panic("NO S FOUND")
}

func expandContent(content [][]rune) (distances [][]int) {
	distances = make([][]int, len(content)*3)
	for i, row := range content {
		for _, tile := range row {
			expandedTile := setMap[tile]
			distances[i*3] = append(distances[i*3], expandedTile[0]...)
			distances[i*3+1] = append(distances[i*3+1], expandedTile[1]...)
			distances[i*3+2] = append(distances[i*3+2], expandedTile[2]...)
		}
	}
	return distances
}

func padContent(content [][]int) [][]int {
	for i, row := range content {
		content[i] = append(append([]int{-2}, row...), -2)
	}
	pad := make([]int, len(content[0]))
	for i := range content[0] {
		pad[i] = -2
	}
	content = append(append([][]int{pad}, content...), pad)
	return content
}

func explore(distances [][]int, x, y int) {
	nextSquare := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	for _, square := range nextSquare {
		if distances[square[0]][square[1]] != -2 {
			if distances[square[0]][square[1]] == -1 || distances[square[0]][square[1]] > distances[x][y]+1 {
				distances[square[0]][square[1]] = distances[x][y] + 1
				explore(distances, square[0], square[1])
			}
		}
	}
}

func findLargest(distances [][]int) (largest int) {
	largest = 0
	for _, row := range distances {
		_max1 := slices.Max(row)
		if _max1 > largest {
			largest = _max1
		}
	}
	return largest
}

func invert(distances [][]int) (inverted [][]int) {
	inverted = make([][]int, len(distances))
	for i, row := range distances {
		for _, tile := range row {
			if tile >= 0 {
				inverted[i] = append(inverted[i], -1)
			} else {
				inverted[i] = append(inverted[i], 0)
			}
		}
	}
	return inverted
}

var count = 1

func scan(distances [][]int) {
	for y := 2; y < len(distances); y += 3 {
		for x := 2; x < len(distances[0]); x += 3 {
			if distances[y][x] == 0 {
				count = 1
				exploreP2(distances, x, y)
			}
		}
	}
}

func exploreP2(distances [][]int, x, y int) bool {
	if !(x-3 >= 0 && y-3 >= 0 && x+3 < len(distances[0]) && y+3 < len(distances)) {
		revert(distances, x, y)
		return false
	}
	distances[y][x] = count
	count++

	nextSquare := [][]int{{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1}, {x, y - 1}, {x, y + 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}}
	for _, square := range nextSquare {
		x = square[0]
		y = square[1]
		if distances[y][x] == 0 {
			if success := exploreP2(distances, x, y); !success {
				return false
			}

		}
	}
	return true
}

func revert(distances [][]int, x, y int) {
	distances[y][x] = -2

	nextSquare := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}

	for _, square := range nextSquare {
		x = square[0]
		y = square[1]
		if !(x >= 0 && y >= 0 && x < len(distances[0]) && y < len(distances)) {
			continue
		}

		if distances[y][x] >= 0 {
			revert(distances, x, y)
		}
	}
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
