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
	n1 := traverse(_map, 1, 1)
	n2 := traverse(_map, 3, 1)
	n3 := traverse(_map, 5, 1)
	n4 := traverse(_map, 7, 1)
	n5 := traverse(_map, 1, 2)
	t := n1 * n2 * n3 * n4 * n5
	fmt.Println(t)
}

func traverse(m [][]rune, r, d int) int {

	boundX := len(m[0])
	boundY := len(m)
	n := 0
	x := 0
	y := 0

	for {	
		
		x = (x + r) % boundX 
		y += d

		if y >= boundY {
			break
		}

		if m[y][x] == '#' {
			n++
		}
 
	}
	
	return n
}
