package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
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
		year, err := strconv.Atoi(val)
		if err == nil && year >= 1920 && year <= 2002 {
			return true
		}
		return false
	},
	"iyr": func(val string) bool {
		year, err := strconv.Atoi(val)
		if err == nil && year >= 2010 && year <= 2020 {
			return true
		}
		return false
	},
	"eyr": func(val string) bool {
		year, err := strconv.Atoi(val)
		if err == nil && year >= 2020 && year <= 2030 {
			return true
		}
		return false
	},
	"hgt": func(val string) bool {
		res, _ := regexp.MatchString("\\d+[cm|in]", val)
		if !res {
			return false
		}
		unit := val[len(val)-2:]
		num, _ := strconv.Atoi(val[:len(val)-2])
		if unit == "cm" {
			return num >= 150 && num <= 193
		} else if unit == "in" {
			return num >= 59 && num <= 76
		}
		return false
	},
	"hcl": func(val string) bool {
		res, _ := regexp.MatchString("^#([0-9a-f]{6})$", val)
		return res
	},
	"ecl": func(val string) bool {
		res, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", val)
		return res
	},
	"pid": func(val string) bool {
		res, _ := regexp.MatchString("^\\d{9}$", val)
		return res
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
		fmt.Println(key, val)
	}
	return true
}

func main() {
	count := 0
	inputLines := getInput()
	for i := 0; i < len(inputLines); i++ {
		doc := getDocumentDetails(inputLines[i])
		if isPassport(doc) {
			count++
			fmt.Println(doc)
		}
	}
	fmt.Println(count)
}
