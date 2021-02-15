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

func parseInstruction(instruction string) (int, int) {
	var command string
	var value int
	fmt.Sscanf(strings.TrimSpace(instruction), "%s %d", &command, &value)
	switch command {
	case "acc":
		return value, 1
	case "jmp":
		return 0, value
	default:
		return 0, 1
	}
}

func main() {
	var executedInstructionMap = map[int]bool{}
	inputLines := getInput()
	var i = 0
	var acc = 0
	_, exists := executedInstructionMap[i]
	for !exists {
		executedInstructionMap[i] = true
		var addToAccumulator, jumpInstructionsBy = parseInstruction(inputLines[i])
		acc += addToAccumulator
		i += jumpInstructionsBy
		_, exists = executedInstructionMap[i]
	}
	fmt.Println(i, acc)
}
