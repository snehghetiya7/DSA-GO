package main

import (
	"fmt"
	"strings"
)

func main() {
	grid := make([][]string, 3)
	var input string
	var word string
	fmt.Scanf("%s", &word)
	wordArr := strings.Split(word, "")
	for i := range grid {
		grid[i] = make([]string, 4)
		for j := range grid[i] {
			fmt.Scanf("%s", &input)
			grid[i][j] = input
		}
	}
	fmt.Println("\033[2J")
	fmt.Println(grid)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == wordArr[j] {
				grid[i][j] = "*"
				break
			}
		}
	}
	fmt.Println(grid)
}
