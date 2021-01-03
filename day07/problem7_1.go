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

func getChildBagList(childBagStr string) ([]string, int) {
	if strings.TrimSpace(childBagStr) == "no other bags." {
		return make([]string, 0), 0
	}
	childBagStrSlice := strings.Split(childBagStr, ",")
	var childBagList []string
	var num int
	var color1, color2 string
	for _, childBagPhrase := range childBagStrSlice {
		fmt.Sscanf(strings.TrimSpace(childBagPhrase), "%d %s %s bag", &num, &color1, &color2)
		childBagList = append(childBagList, fmt.Sprintf("%s %s", color1, color2))
	}
	return childBagList, num
}

func getUniqueParents(childToParent map[string]map[string]int, targetColor string, alreadySeen map[string]bool) int {
	parentMap, _ := childToParent[targetColor]
	recursiveCount := 0
	if parentMap != nil {
		for parentBag := range parentMap {
			if _, exists := alreadySeen[parentBag]; !exists {
				recursiveCount++
				alreadySeen[parentBag] = true
				recursiveCount += getUniqueParents(childToParent, parentBag, alreadySeen)
			}
		}
	}
	return recursiveCount
}

func main() {
	inputLines := getInput()
	// Map to record every parent bag that contains a
	// specific child bag, like a reverse tree
	childToParent := make(map[string]map[string]int)

	// Record each bag rule in the inpiut and construct
	// the above declared map
	for _, bagRule := range inputLines {

		splitBagRule := strings.Split(bagRule, "bags contain")
		parentBag := strings.TrimSpace(splitBagRule[0])
		childBagList, _ := getChildBagList(splitBagRule[1])
		for _, childBag := range childBagList {
			parentMap, exists := childToParent[childBag]
			if !exists {
				parentMap = make(map[string]int)
				childToParent[childBag] = parentMap
			}

			parentMap[parentBag] = 1
		}
	}

	// fmt.Println(childToParent["shiny gold"])

	// After construction of map, we will traverse the
	// map recursively to see how many bags can have a
	// shiny gold bag
	targetBag := "shiny gold"
	// fmt.Println(childToParent)
	numOfParentBags := getUniqueParents(childToParent, targetBag, make(map[string]bool))
	fmt.Println(numOfParentBags)

}
