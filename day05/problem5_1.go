package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
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

func getSeatID(boardingPassNo string) int {
	temp := strings.ReplaceAll(boardingPassNo, "F", "0")
	temp2 := strings.ReplaceAll(temp, "B", "1")
	temp3 := strings.ReplaceAll(temp2, "L", "0")
	binaryStr := strings.ReplaceAll(temp3, "R", "1")
	seatID, _ := strconv.ParseInt(binaryStr, 2, 64)
	return int(seatID)
}

func main() {
	var largestSeatID int
	inputLines := getInput()
	for i := 0; i < len(inputLines); i++ {
		seatID := getSeatID(inputLines[i])
		if seatID > largestSeatID {
			largestSeatID = seatID
		}
	}
	fmt.Println(largestSeatID)
}
