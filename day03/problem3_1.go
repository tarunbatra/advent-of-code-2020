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

// TREE is a sign representing a tree in input
const TREE = "#"

func getNextPosition(line string, rowIndex int) string {
	columnIndex := rowIndex * 3
	return string(line[columnIndex%len(line)])
}

func main() {
	count := 0
	inputLines := getInput()
	for i := 1; i < len(inputLines); i++ {
		nextPos := getNextPosition(inputLines[i], i)
		// fmt.Println(nextPos)
		if nextPos == TREE {
			count++
		}
	}
	fmt.Println(count)
}
