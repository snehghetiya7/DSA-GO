package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	var length int
	var sum int
	for i := 0; i < input; i++ {
		fmt.Scanf("%d", &length)
		var arr []int
		for j := 0; j < length; j++ {
			var temp int
			fmt.Scanf("%d", &temp)
			arr = append(arr, temp)
		}
		fmt.Scanf("%d", &sum)
		findIndex(arr, sum)
	}
}

func findIndex(arr []int, sum int) {
	for i := 0; i < len(arr); i++ {
		sum = sum - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if sum == arr[j] {
				println(i, j)
			}
		}
	}
}
