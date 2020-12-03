package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	min int
	max int
	letter string 

}

func (p *Policy) IsValid(s string) bool {
	c := strings.Count(s, p.letter)
	if c >= p.min && c <= p.max {
		return true
	}
	return false
}

func Parse(s string) (string, Policy) {
	var min, max int
	var pw, l string

	split := strings.Split(s, " ")
	mm := strings.Split(split[0], "-")
	min, _ = strconv.Atoi(mm[0])
	max, _ = strconv.Atoi(mm[1])
	l = strings.TrimRight(split[1], ":") 
	pw = split[2]

	return pw, Policy { min, max, l }
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