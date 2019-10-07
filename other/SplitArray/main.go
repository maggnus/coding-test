package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 6, 4, 3, 3, 2, 4, 7, 3, 2, 5, 2}
	fmt.Println(arr)
	arrLeft := make([]int, len(arr))
	arrRight := make([]int, len(arr))
	var sumLeft, sumRight int
	lenght := len(arr)
	for k, v := range arr {
		sumLeft = sumLeft + v
		sumRight = sumRight + arr[lenght-k-1]
		arrLeft[k] = sumLeft
		arrRight[lenght-k-1] = sumRight
	}
	for k := range arr {
		if arrLeft[k] == arrRight[k] {
			fmt.Println(arr[:k])
			fmt.Println(arr[k:])
		}
	}
}
