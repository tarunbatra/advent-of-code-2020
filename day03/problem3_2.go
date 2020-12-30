package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput() []string {
	filePath, _ := filepath.Abs("./input.txt")
	input, err := ioutil.ReadFile(filePath)
	check(err)
	return strings.Split(string(input), "\n")
}

func getNextPosition(line string, rowIndex int, stepRight int) string {
	columnIndex := rowIndex * stepRight
	return string(line[columnIndex%len(line)])
}

func main() {
	inputLines := getInput()
	// 5 algos to traverse the slope. {3, 1} indicates, stepRight of 3 and stepDown of 1
	algos := [5][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	res := 1
	for i := 0; i < len(algos); i++ {
		result := traverseSlope(inputLines, algos[i][1], algos[i][0])
		fmt.Printf("If we step right %d and step down %d, we will encounter %d trees.\n", algos[i][1], algos[i][0], result)
		res *= result
	}
	fmt.Printf("Result of multiplication of all the trees encountered is %d\n", res)
}

func traverseSlope(geo []string, stepDown int, stepRight int) int {
	const TREE = "#"
	count := 0
	step := 1
	for i := stepDown; i < len(geo); i += stepDown {
		nextPos := getNextPosition(geo[i], step, stepRight)
		if nextPos == TREE {
			count++
		}
		step++
	}
	return count
}
