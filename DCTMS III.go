package main

import "fmt"

func canPlace(words []string, p, q int) bool {
	for i, n := 0, len(words); i < n; i++ {
		if words[i][q] < words[i][p] {
			return false
		}
	}
	return true
}

func intMax2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dctms(i int, words []string, mem []int) int {
	if i == len(words[0]) {
		return 0
	}
	if mem[i] != -1 {
		return mem[i]
	}
	max := 0
	for j, m := i+1, len(words[0]); j < m; j++ {
		if canPlace(words, i, j) {
			tmp := dctms(j, words, mem)
			max = intMax2(max, tmp)
		}
	}
	mem[i] = max + 1
	return mem[i]
}

func dctms3(words []string) int {
	mem := make([]int, len(words[0]))
	for i := range mem {
		mem[i] = -1
	}
	res := 0
	for i, m := 0, len(words[0]); i < m; i++ {
		res = intMax2(res, dctms(i, words, mem))
	}
	return res
}

func main() {
	words := []string{
		"babca",
		"bbazb",
	}
	fmt.Println(len(words[0]) - dctms3(words))
}
