package main

import (
	"fmt"
)

const (
	mod   = 1000000007
	prime = 16008019
)

var po []int64
var hash []int64

func CalcHash(s string) {
	n := len(s)
	po = make([]int64, n)
	hash = make([]int64, n)
	po[0] = 1
	for i := 1; i < n; i++ {
		po[i] = (po[i-1] * prime) % mod
	}
	hash[0] = int64(s[0])
	for i := 1; i < n; i++ {
		hash[i] = (hash[i-1]*prime + int64(s[i])) % mod
	}
}

func RangeHash(l, r int) int64 {
	if l == 0 {
		return hash[r]
	}
	return ((hash[r]-hash[l-1]*po[r-l+1])%mod + mod) % mod
}

var mem []int

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func call(i, n int) int {
	if i == n/2 {
		return 1
	}
	if mem[i] != -1 {
		return mem[i]
	}

	max := 1
	for j := i; j < n/2; j++ {
		p := n - j - 1
		q := n - i - 1
		if RangeHash(i, j) == RangeHash(p, q) {
			tmp := 2 + call(j+1, n)
			max = intMax(max, tmp)
		}
	}
	mem[i] = max
	return max
}

func LCPD(s string) int {
	mem = make([]int, len(s)+1)
	for i := range mem {
		mem[i] = -1
	}
	CalcHash(s)

	return call(0, len(s))
}

func main() {
	s := "abcde"
	fmt.Println(LCPD(s))
}
