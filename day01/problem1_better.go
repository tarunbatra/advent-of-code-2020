package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loopAndAdd(list []int, startPoint int, sum int, level int) int {
	mult := -1
	for i := startPoint; i < len(list); i++ {
		newSum := sum + list[i]
		if newSum > 2020 {
			continue
		}
		if level < 2 {
			ret := loopAndAdd(list, i+1, newSum, level+1)
			if ret != -1 {
				mult = list[i] * ret
				break
			}
		} else if newSum == 2020 {
			mult = list[i]
			break
		}
	}
	return mult
}

func main() {

	input, err := ioutil.ReadFile("/Users/tarunbatra/Personal/advent-of-code-2020/day01/input.txt")
	check(err)
	inputLines := strings.Split(string(input), "\n")
	nums := make([]int, len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		num, err := strconv.Atoi(inputLines[i])
		check(err)
		nums[i] = num
	}
	ret := loopAndAdd(nums, 0, 0, 0)

	if ret != 267520550 {
		fmt.Println(ret)
		panic(errors.New("code failed. expected 267520550"))
	}
}
