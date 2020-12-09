package main

import (
	"fmt"
	"os"
	"bufio"
)

var _map [][]rune

func init() {
	_map = [][]rune{}
	 
	file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
	
	scanner := bufio.NewScanner(file)

    for scanner.Scan() {
		slice := []rune(scanner.Text())
		_map = append(_map, slice)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
	}
}

func main() {
	xl := len(_map[0])
	x := 3
	sum := 0
	// skip first row
	for _, row := range _map[1:] {
		if row[x] == '#' {
			sum++
		}
		x = (x + 3) % xl
	}
	fmt.Println(sum)
}
