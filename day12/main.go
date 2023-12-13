package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var checkRegex = regexp.MustCompile("#+")
var optionalRegex = regexp.MustCompile("(\\?|#)+")

func main() {
	fmt.Println("PART 1 Test:")
	part1("test.txt")
	fmt.Println("\nPART 2 Test:")
	// part2("test.txt")

	fmt.Println("\nPART 1:")
	part1("input.txt")
	fmt.Println("\nPART 2:")
	part2("input.txt")
}

func regexForTheWin(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)

	code1 := []string{}
	code2 := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		code1 = append(code1, regexp.MustCompile("(\\?|#|\\.)+").FindString(line))
		code2 = append(code2, regexp.MustCompile("\\d+").FindAllString(line, -1))
	}
	file.Close()

	for i, c := range code1 {
		for j := 0; j < 4; j++ {
			code1[i] = code1[i] + "?" + c
		}
	}
	for i, c := range code2 {
		for j := 0; j < 4; j++ {
			code2[i] = append(code2[i], c...)
		}
	}

	code2Num := make([][]int, len(code2))

	for i := range code2 {
		for _, c := range code2[i] {
			_c, err := strconv.Atoi(c)
			check(err)
			code2Num[i] = append(code2Num[i], _c)
		}
	}
	sum := 0
	for i, s := range code1 {
		code1[i] = "." + s + "."
		regexStr := "(?="
		for _, val := range code2Num[i] {
			v := strconv.Itoa(val)
			regexStr += "(\\?|#){" + v + "}(\\?|\\.)+"
		}
		regexStr += ")"
		secretRegEx := regexp.MustCompile(regexStr)
		sum += len(secretRegEx.FindAllStringIndex("."+s+".", -1))
		println(sum)
	}

}

func part1(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)

	code1 := []string{}
	code2 := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		code1 = append(code1, regexp.MustCompile("(\\?|#|\\.)+").FindString(line))
		code2 = append(code2, regexp.MustCompile("\\d+").FindAllString(line, -1))
	}
	file.Close()

	code2Num := make([][]int, len(code2))

	for i := range code2 {
		for _, c := range code2[i] {
			_c, err := strconv.Atoi(c)
			check(err)
			code2Num[i] = append(code2Num[i], _c)
		}
	}

	sum := 0
	for i, c1 := range code1 {
		if satisfiesInput(c1, code2Num[i]) {
			sum++
			continue
		}
		qLocs := regexp.MustCompile("\\?").FindAllStringIndex(c1, -1)
		for j := 0; j < int(math.Pow(2, float64(len(qLocs)))); j++ {
			codeCopy := []rune(c1)
			jCopy := j
			for _, loc := range qLocs {
				if (jCopy & 1) == 1 {
					codeCopy[loc[0]] = '#'
				}
				jCopy = jCopy >> 1
			}
			satisfaction := satisfiesInput(string(codeCopy), code2Num[i])
			if satisfaction {
				sum++
			}
		}

	}
	println(sum)
}

func satisfiesInput(code1 string, code2 []int) bool {
	indexes := checkRegex.FindAllStringIndex(code1, -1)
	if len(indexes) != len(code2) {
		return false
	}
	for i, index := range indexes {
		if index[1]-index[0] != code2[i] {
			return false
		}
	}
	return true
}

func part2(filename string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)

	code1 := []string{}
	code2 := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		code1 = append(code1, regexp.MustCompile("(\\?|#|\\.)+").FindString(line))
		code2 = append(code2, regexp.MustCompile("\\d+").FindAllString(line, -1))
	}
	file.Close()

	for i, c := range code1 {
		for j := 0; j < 4; j++ {
			code1[i] = code1[i] + "?" + c
		}
	}
	for i, c := range code2 {
		for j := 0; j < 4; j++ {
			code2[i] = append(code2[i], c...)
		}
	}

	code2Num := make([][]int, len(code2))

	for i := range code2 {
		for _, c := range code2[i] {
			_c, err := strconv.Atoi(c)
			check(err)
			code2Num[i] = append(code2Num[i], _c)
		}
	}

	codeSplit := make([][]string, len(code1))
	for i, c := range code1 {
		codeSplit[i] = append(codeSplit[i], regexp.MustCompile("(\\?|#)+").FindAllString(c, -1)...)
	}
	pad(codeSplit)

	sum := 0
	for i, line := range codeSplit {

		sum += solveLine(line, code2Num[i])
		if i%1 == 0 {
			fmt.Printf("Line %d: %d\n", i+1, sum)
		}
	}

	println(sum)
}

func solveLine(snippets []string, numbers []int) (count int) {
	sum1 := 0
	sum2 := 0
	sum3 := 0
	for _, s := range snippets {
		sum1 += strings.Count(s, "#")
	}
	for _, n := range numbers {
		sum2 += n
	}
	for _, s := range snippets {
		sum3 += strings.Count(s, "?")
	}

	if sum1+sum3 < sum2 {
		return 0
	}

	if len(snippets) == 0 {
		if len(numbers) == 0 {
			return 1
		}
		return 0
	}
	current := snippets[0]

	if sum1 == sum2 {
		current = strings.ReplaceAll(current, "?", ".")
	}

	if strings.Count(current, "?") == 0 {
		groups := regexp.MustCompile("#+").FindAllString(snippets[0], -1)
		for i := range groups {
			if len(numbers) == 0 {
				numbers = append(numbers, 0)
			}
			if len(groups[i]) != numbers[i] {
				return 0
			}
		}
		return solveLine(snippets[1:], numbers[len(groups):])
	}

	solutionTally := solutionsPerSnippet(current, numbers)
	count = 0
	for _, solution := range maps.Keys(solutionTally) {
		if len(snippets) <= 1 && solution == 0 {
			count += solutionTally[solution]
			continue
		} else if len(snippets) <= 1 {
			continue
		}
		diff := solutionTally[solution] * solveLine(snippets[1:], numbers[len(numbers)-solution:])
		count += diff
	}
	return count
}

func pad(ss [][]string) {
	for _, s := range ss {
		for i := range s {
			s[i] = "." + s[i] + "."
		}
	}
}

func solutionsPerSnippet(snippet string, numbers []int) map[int]int {
	currentCount := numbers[:]
	for i := 0; i < len(snippet); {
		if snippet[i] == '#' {
			if len(currentCount) == 0 {
				return map[int]int{}
			}
			for j := 0; j < currentCount[0]; j++ {
				if snippet[i+j] == '?' {
					snippet = snippet[:i+j] + "#" + snippet[i+j+1:]
				}
				if snippet[i+j] == '.' {
					return make(map[int]int)
				}
			}
			if snippet[i+currentCount[0]] == '#' {
				return make(map[int]int)
			}
			if snippet[i+currentCount[0]] == '?' {
				snippet = snippet[:i+currentCount[0]] + "." + snippet[i+currentCount[0]+1:]
			}

			i += currentCount[0]
			currentCount = currentCount[1:]
			continue
		}
		if snippet[i] == '?' {
			if len(currentCount) == 0 {
				return solutionsPerSnippet(snippet[i+1:], currentCount)
			}
			streak := 1
			for snippet[i+streak] == '?' {
				streak++
			}
			if snippet[i+streak] == '.' {
				diffGroups := streak + 1
				j := 0
				rMap := solutionsPerSnippet(snippet[i+streak:], currentCount)
				for {
					j++
					if len(currentCount) < j {
						break
					}
					diffGroups -= currentCount[j-1]

					sols := map[int]int{}
					combs := 0
					if diffGroups < j {
						break
					}
					if len(currentCount) == j {
						sols = solutionsPerSnippet(snippet[i+streak:], []int{})
					} else {
						sols = solutionsPerSnippet(snippet[i+streak:], currentCount[j:])
					}
					sols = solutionsPerSnippet(snippet[i+streak:], currentCount[j:])
					combs = choose(diffGroups, j)

					for _, k := range maps.Keys(sols) {
						rMap[k] += sols[k] * combs
					}

				}
				return rMap

			}
			s2 := map[int]int{}
			if len(currentCount) != 0 {
				c2 := "#" + snippet[i+1:]
				s2 = solutionsPerSnippet(c2, currentCount)
			}
			c1 := snippet[i+1:]
			s1 := solutionsPerSnippet(c1, currentCount)
			for _, v := range maps.Keys(s1) {
				s2[v] += s1[v]
			}
			return s2
		}
		i++
	}
	return map[int]int{len(currentCount): 1}
}

func checkValueSubstring(substrings []string, values []int) (status int) {
	if len(substrings) == 0 {
		if len(values) == 0 {
			return 1
		}
		return -1
	}
	sum1 := 0
	sum2 := 0
	sum3 := 0
	for _, s := range substrings {
		sum1 += strings.Count(s, "#")
	}
	for _, n := range values {
		sum2 += n
	}
	for _, s := range substrings {
		sum3 += strings.Count(s, "?")
	}
	if sum1+sum3 < sum2 {
		return -1
	}
	return 0
}

func choose(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sum(s []int) (tot int) {
	for _, e := range s {
		tot += e
	}
	return tot
}
