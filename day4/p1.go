package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
)

type passport map[string]string

func isValid(p passport, f []string) bool {
	for _, field := range f {
		if _, ok := p[field]; !ok {
			return false
		}
	}
	return true
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

	data := strings.Split(string(file), "\r\n\r\n")
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
	n := 0

	for _, s := range data {
		p := newPassport(s)
		if isValid(p, fields) {
			n++
		}
	}
	
	fmt.Printf("%d valid passports", n)
}