package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	first int
	second int
	char rune 
}

func (p *Policy) IsValid(s string) bool {
	str := []rune(s)  
	
	if str[p.first] == p.char && str[p.second] != p.char {
		return true
	}

	if str[p.first] != p.char && str[p.second] == p.char {
		return true
	}

	return false
}

func Parse(s string) (string, Policy) {
	var first, second int
	var pw string
	var char rune

	split := strings.Split(s, " ")
	mm := strings.Split(split[0], "-")
	first, _ = strconv.Atoi(mm[0])
	second, _ = strconv.Atoi(mm[1]) 
	runes := []rune(split[1])
	char = runes[0]
	pw = split[2]

	return pw, Policy { first - 1, second - 1, char }
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
	
	scanner := bufio.NewScanner(file)
	valid := 0


    for scanner.Scan() {
		pw, policy := Parse(scanner.Text())
		if policy.IsValid(pw) {
			valid += 1
		}
    }
	
	fmt.Println(valid)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
	}

}