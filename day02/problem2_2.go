package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput() []string {
	input, err := ioutil.ReadFile("/Users/tarunbatra/Personal/advent-of-code-2020/day02/input.txt")
	check(err)
	return strings.Split(string(input), "\n")
}

func parseInputLine(line string) (int, int, string, string) {
	split := strings.Split(line, ": ")
	leftSplit := strings.Split(split[0], " ")
	digitSplit := strings.Split(leftSplit[0], "-")
	start, err1 := strconv.Atoi(digitSplit[0])
	check(err1)
	end, err2 := strconv.Atoi(digitSplit[1])
	check(err2)
	return start, end, leftSplit[1], split[1]
}

func createPattern(start int, end int, char string) *regexp.Regexp {
	r, _ := regexp.Compile(fmt.Sprintf("%s", char))
	return r
}

func didPasswordMatch(matchIndices [][]int, start int, end int) bool {
	startMatch := false
	endMatch := false
	for i := 0; i < len(matchIndices); i++ {
		currentIndex := matchIndices[i][0]
		if currentIndex > end {
			break
		}
		if currentIndex == start-1 {
			startMatch = true
		} else if currentIndex == end-1 {
			endMatch = true
		}
	}
	return startMatch && !endMatch || endMatch && !startMatch
}

func main() {
	count := 0
	inputLines := getInput()
	for i := 0; i < len(inputLines); i++ {
		start, end, char, str := parseInputLine(inputLines[i])
		pattern := createPattern(start, end, char)
		matchCount := pattern.FindAllIndex([]byte(str), -1)
		if didPasswordMatch(matchCount, start, end) {
			count++
		} else {
			// fmt.Println(inputLines[i])
		}
	}
	fmt.Println(count)
}
