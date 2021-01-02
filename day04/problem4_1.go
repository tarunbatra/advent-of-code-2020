package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
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

func getDocumentDetails(inputStr string) map[string]string {
	doc := make(map[string]string)
	split := regexp.MustCompile("[ \n]+").Split(inputStr, -1)
	for i := 0; i < len(split); i++ {
		fields := strings.Split(split[i], ":")
		doc[fields[0]] = fields[1]
	}
	return doc
}

var validFields = map[string]func(string) bool{
	"byr": func(val string) bool {
		return true
	},
	"iyr": func(val string) bool {
		return true
	},
	"eyr": func(val string) bool {
		return true
	},
	"hgt": func(val string) bool {
		return true
	},
	"hcl": func(val string) bool {
		return true
	},
	"ecl": func(val string) bool {
		return true
	},
	"pid": func(val string) bool {
		return true
	},
}

func isPassport(doc map[string]string) bool {
	for key, validationFunc := range validFields {
		val, exist := doc[key]
		if !exist {
			return false
		}
		if !validationFunc(val) {
			return false
		}
	}

	return true
}

func main() {
	count := 0
	inputLines := getInput()
	for i := 0; i < len(inputLines); i++ {
		doc := getDocumentDetails(inputLines[i])
		// fmt.Println(doc)
		if isPassport(doc) {
			count++
		}
	}
	fmt.Println(count)
}
