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

func getChildBagMap(childBagStr string) map[string]int {
	if strings.TrimSpace(childBagStr) == "no other bags." {
		return make(map[string]int, 0)
	}
	childBagStrSlice := strings.Split(childBagStr, ",")
	var childBagMap = map[string]int{}
	var num int
	var color1, color2 string
	for _, childBagPhrase := range childBagStrSlice {
		fmt.Sscanf(strings.TrimSpace(childBagPhrase), "%d %s %s bag", &num, &color1, &color2)
		color := fmt.Sprintf("%s %s", color1, color2)
		childBagMap[color] = num
	}
	return childBagMap
}

func getNumberOfChildBags(parentToChild map[string]map[string]int, targetColor string, alreadySeen map[string]bool) int {
	childMap := parentToChild[targetColor]
	fmt.Println(targetColor, childMap)
	recursiveCount := 0
	if childMap != nil {
		for childBag, childCount := range childMap {
			recursiveCount += childCount + childCount*getNumberOfChildBags(parentToChild, childBag, alreadySeen)
		}
	}
	return recursiveCount
}

func main() {
	inputLines := getInput()
	// Map to record every child bag inside a
	// specific parent bag, like a reverse tree
	parentToChild := make(map[string]map[string]int)

	// Record each bag rule in the inpiut and construct
	// the above declared map
	for _, bagRule := range inputLines {

		splitBagRule := strings.Split(bagRule, "bags contain")
		parentBag := strings.TrimSpace(splitBagRule[0])
		childMap, exists := parentToChild[parentBag]
		if !exists {
			childMap = make(map[string]int)
			parentToChild[parentBag] = childMap
		}
		childBagMap := getChildBagMap(splitBagRule[1])
		for color, count := range childBagMap {
			childMap[color] = count
		}
	}

	// After construction of map, we will traverse the
	// map recursively to see how many bags can have a
	// shiny gold bag
	targetBag := "shiny gold"
	numberOfChildBags := getNumberOfChildBags(parentToChild, targetBag, make(map[string]bool))
	fmt.Println(numberOfChildBags)

}
