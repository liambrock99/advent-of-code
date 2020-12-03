package main
 
import (
    "fmt"
    "os"
	"bufio"
	"strconv"
	"strings"
)
 
func find2(s []int) int {
	for _, x := range(s) {
		for _, y := range(s) {
			if x + y == 2020 {
				return x * y
			}
		}
	}
	return 0
}

func find3(s []int) int {
	for _, x := range(s) {
		for _, y := range(s) {
			for _, z := range(s) {
				if x + y + z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return 0
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
	
	nums := make([]int, 0)
	scanner := bufio.NewScanner(file)
	
    for scanner.Scan() {
		str := scanner.Text()
		str = strings.TrimSpace(str)
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, num)
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
	}

	fmt.Println(find2(nums))
	fmt.Println(find3(nums))
}