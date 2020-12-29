package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	result := 2020

	input, err := ioutil.ReadFile("/Users/tarunbatra/Personal/advent-of-code-2020/day01/input.txt")
	check(err)
	inputLines := strings.Split(string(input), "\n")
	nums := make([]int, len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		num, err := strconv.Atoi(inputLines[i])
		check(err)
		nums[i] = num
	}
	var ret int
	for i := 0; i < len(inputLines); i++ {
		addresult := nums[i]
		if addresult > result {
			continue
		}
		for j := i + 1; j < len(inputLines); j++ {
			addaddresult := addresult + nums[j]
			if addaddresult > result {
				continue
			}
			for k := j + 1; k < len(inputLines); k++ {
				addaddaddresult := addaddresult + nums[k]
				if addaddaddresult == result {
					// fmt.Println(addaddaddresult-addaddresult, addaddresult-addresult, addresult)
					ret = nums[i] * nums[j] * nums[k]
					// fmt.Println(i, j, k)
					// fmt.Println(nums[i], nums[j], nums[k])
					break
				}
			}
		}
	}
	if ret != 267520550 {
		panic("code failed")
	} else {
	}
}
