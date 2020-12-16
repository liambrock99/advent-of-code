package main

import (
	"fmt"
	"log"
	"strings"
	"io/ioutil"
)


func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(file), "\r\n")
	highest := 0
	for _, s := range input {
		sid := decode([]rune(s))
		if sid > highest {
			highest = sid
		}
	}
	fmt.Printf("The highest seat ID is %d", highest)
}

// returns new upper bound
func upper(l, u int) int {
	return u - ((u - l) / 2) - 1
}

// returns new lower bound
func lower (l, u int) int {
	return l + ((u - l) / 2) + 1
}

func decode(runes []rune) int {
	rowL, rowU := 0, 127 
	colL, colU := 0, 7
	for _, r := range runes {
		switch r {
		case 'F':
			rowU = upper(rowL, rowU) // take lower half (new upper bound)
		case 'B':
			rowL = lower(rowL, rowU) // take upper half (new lower bound)
		case 'L':
			colU = upper(colL, colU)
		case 'R':
			colL = lower(colL, colU)
		}
	}
	return rowL * 8 + colL
}