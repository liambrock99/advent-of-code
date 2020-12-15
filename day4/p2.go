package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
	"log"
)

type passport map[string]string



func isValid(pp passport, required []string) bool {
	for _, field := range required {
		val, ok := pp[field]
		if !ok {
			return false
		}
		switch field {
		case "byr":
			if ok := validYear(val, 1920, 2002); !ok {
				return false
			}
		case "iyr":
			if ok := validYear(val, 2010, 2020); !ok {
				return false
			}
		case "eyr":
			if ok := validYear(val, 2020, 2030); !ok {
				return false
			}
		case "hgt":
			if ok := hgt(val); !ok {
				return false
			}
		case "hcl":
			if ok := hcl(val); !ok {
				return false
			}
		case "ecl":
			if ok := ecl(val); !ok {
				return false
			}
		case "pid":
			if ok := pid(val); !ok {
				return false
			}
		}
	}
	return true
}

func validYear(s string, min, max int) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	if i >= min && i <= max {
		return true
	}
	return false
}

func hgt(s string) bool {
	runes := []rune(s)
	val, err := strconv.Atoi(string(runes[:len(runes)-2]))
	if err != nil {
		log.Fatal()
	}
	unit := string(runes[len(runes)-2:])

	if unit == "cm" {
		if val >= 150 && val <= 193 {
			return true
		}
	}

	if unit == "in" {
		if val >= 59 && val <= 76 {
			return true
		}
	}

	return false
}

func hcl(s string) bool {
	match, err := regexp.MatchString(`^#[0-9a-f]{6}$`, s)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

func ecl(s string) bool {
	valid := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}
	for _, v := range valid {
		if s == v {
			return true
		}
	}
	return false
}

func pid(s string) bool {
	match, err := regexp.MatchString(`^[0-9]{9}$`, s)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

func newPassport(s string) passport {
	passport := make(passport)
	s = strings.ReplaceAll(s, "\r\n", " ")
	fields := strings.Fields(s)
	for _, s := range fields {
		kv := strings.Split(s, ":")
		passport[kv[0]] = kv[1]
	}
	return passport
}

func main() {	 
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(file), "\r\n\r\n")
	valid := 0
	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	for _, s := range input {
		pp := newPassport(s)
		if isValid(pp, fields) {
			valid++
		}
	}	
	fmt.Println(valid)
}