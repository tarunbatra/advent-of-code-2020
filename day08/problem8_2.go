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

func parseInstructions(instruction string) (int, int, string) {
	var command string
	var value int
	fmt.Sscanf(strings.TrimSpace(instruction), "%s %d", &command, &value)
	switch command {
	case "acc":
		return value, 1, command
	case "jmp":
		return 0, value, command
	default:
		return 0, 1, command
	}
}

var instructionSwap = map[string]string{
	"jmp": "nop",
	"nop": "jmp",
}

func executeTillTerminates(currentInstruction string, allInstructions []string, current int, acc int, executedInstructionMap map[int]bool) (bool, int) {
	seen, _ := executedInstructionMap[current]
	if seen {
		return false, acc
	}
	// If reached end, return true
	var addToAccumulator, jumpInstructionsBy, command = parseInstructions(currentInstruction)
	if current+jumpInstructionsBy == len(allInstructions) {
		return true, acc + addToAccumulator
	}
	// If neither already seen nor reached end, execute the next instructions
	executedInstructionMap[current] = true
	var treeResult, treeAccValue = executeTillTerminates(allInstructions[current+jumpInstructionsBy], allInstructions, current+jumpInstructionsBy, acc+addToAccumulator, executedInstructionMap)
	if treeResult {
		return treeResult, treeAccValue
	}
	// If no other instruction succeeds, try swapping the command
	if replaceCommand, exists := instructionSwap[command]; exists {
		executedInstructionMap[current] = false
		mutatedInstruction := strings.Replace(currentInstruction, command, replaceCommand, 1)
		mutatedExecutedInstructionMap := make(map[int]bool)
		for k, v := range executedInstructionMap {
			mutatedExecutedInstructionMap[k] = v
		}
		return executeTillTerminates(mutatedInstruction, allInstructions, current, acc, mutatedExecutedInstructionMap)
	}
	return false, acc
}

func main() {
	var executedInstructionMap = map[int]bool{}
	inputLines := getInput()
	_, acc := executeTillTerminates(inputLines[0], inputLines, 0, 0, executedInstructionMap)
	fmt.Println(acc)
}
