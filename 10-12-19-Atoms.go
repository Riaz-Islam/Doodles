package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func getName(s *string, idx *int) string {
	name := ""
	n := len(*s)
	for *idx < n && unicode.IsLetter(rune((*s)[*idx])) && !(len(name) > 0 && unicode.IsUpper(rune((*s)[*idx]))) {
		name += string((*s)[*idx])
		(*idx)++
	}
	return name
}

func getNum(s *string, idx *int) int {
	num := ""
	n := len(*s)
	for *idx < n && unicode.IsDigit(rune((*s)[*idx])) {
		num += string((*s)[*idx])
		(*idx)++
	}
	val, err := strconv.Atoi(num)
	if err != nil {
		val = 1
	}
	return val
}

func calc(s *string, idx *int) map[string]int {
	cnt := map[string]int{}
	n := len(*s)
	for *idx < n && (*s)[*idx] != ')' {
		if (*s)[*idx] == '(' {
			(*idx)++
			bucket := calc(s, idx)
			c := getNum(s, idx)
			for k, v := range bucket {
				cnt[k] += v * c
			}
		} else {
			name := getName(s, idx)
			c := getNum(s, idx)
			cnt[name] += c
		}
	}
	(*idx)++
	return cnt
}

func convToString(m map[string]int) string {
	type elem struct {
		name  string
		count int
	}
	e := []elem{}
	for k, v := range m {
		e = append(e, elem{
			name:  k,
			count: v,
		})
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].name < e[j].name
	})
	res := ""
	for _, el := range e {
		if el.count > 0 {
			res += el.name
		}
		if el.count > 1 {
			res += strconv.Itoa(el.count)
		}
	}
	return res
}

func atoms(s string) string {
	m := calc(&s, new(int))
	return convToString(m)
}

func main() {
	//s := "Ca(CO2)10(Br24(MnO4)100)5HgK2Cr2O7"
	s := "((K(OH)2)5Na)10PSO4"
	fmt.Println(atoms(s))
}
