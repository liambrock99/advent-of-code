package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strconv"
	"strings"
)

type Policy struct {
	first int
	second int
	letter byte 

}

func (p *Policy) IsValid(s string) bool {
	c := strings.Count(s, p.letter)
	if c >= p.min && c <= p.max {
		return true
	}
	return false
}

func main() {
    // file, err := os.Open("input.txt")
    // if err != nil {
    //     fmt.Println(err)
    // }
    // defer file.Close()
	
	// scanner := bufio.NewScanner(file)

    // for scanner.Scan() {
	
    // }
	
    // if err := scanner.Err(); err != nil {
    //     fmt.Println(err)
	// }

	s := "axczz"
	fmt.Println(strings.IndexByte(s, 'z'))
	fmt.Println(strings.LastIndexByte(s, 'z'))
}