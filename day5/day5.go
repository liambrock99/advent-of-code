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
	existing := []int{}
	possible := []int{}
	max := 127 * 8 + 7

	for _, s := range input {
		sid := decode([]rune(s))
		existing = append(existing, sid)
	}
 
	for i := 0; i <= max; i++ {
		// see if i matches any input SID
		if !exists(i, existing) {
			// see if SID i + 1 and i - 1 exist
			if exists(i + 1, existing) && exists(i - 1, existing) {
				possible = append(possible, i)
			}
		}
	}

	fmt.Printf("Yoru boarding pass is %d\n", possible)

		

	highest := 0
	for _, sid := range existing {
		if sid > highest {
			highest = sid
		}
	}
	fmt.Printf("The highest seat ID is %d\n", highest)
}

func exists(i int, ints []int) bool {
	for _, n := range ints {
		if i == n {
			return true
		}
	}
	return false
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