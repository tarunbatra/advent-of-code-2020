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
	return strings.Split(string(input), "\n\n")
}

func main() {
	inputLines := getInput()
	totalCount := 0
	for _, group := range inputLines {
		groupMap := make(map[string]int)
		passengerList := strings.Split(string(group), "\n")
		for _, answers := range passengerList {
			for _, answer := range answers {
				if _, exists := groupMap[string(answer)]; exists {
					groupMap[string(answer)]++
				} else {
					groupMap[string(answer)] = 1
				}
			}
		}
		totalCount += len(groupMap)
	}
	fmt.Println(totalCount)
}
