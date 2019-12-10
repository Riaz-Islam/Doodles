package main

import (
	"fmt"
	"math"
)

func mctms(arr []int) int {
	suffixMin := make([]int, len(arr)+1)
	suffixMin[len(arr)] = math.MaxInt32
	for i := len(arr) - 1; i >= 0; i-- {
		suffixMin[i] = suffixMin[i+1]
		if arr[i] < suffixMin[i] {
			suffixMin[i] = arr[i]
		}
	}

	max := math.MinInt32
	res := 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
		if max <= suffixMin[i+1] {
			max = math.MinInt32
			res++
		}
	}
	return res
}

func main() {
	arr := []int{2, 1, 3, 3, 4, 5}
	fmt.Println(mctms(arr))
}
