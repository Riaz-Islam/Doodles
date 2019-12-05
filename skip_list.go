package main

import (
	"fmt"
	"math"
	"math/rand"
)

type SkipNode struct {
	value int
	level int
	right *SkipNode
	down  *SkipNode
}

type SkipList struct {
	topLeft  *SkipNode
	maxLevel int
}

func (s *SkipList) GenerateRandomLevel() int {
	l := 0
	for l < s.maxLevel {
		if rand.Intn(2) == 0 {
			l++
		} else {
			return l
		}
	}
	return l
}

func Constructor() *SkipList {
	//rand.Seed(time.Now().UnixNano())
	s := &SkipList{}
	s.maxLevel = 16
	s.topLeft = &SkipNode{
		value: math.MinInt32,
		level: s.maxLevel,
	}
	cur := s.topLeft
	for i := s.maxLevel - 1; i >= 0; i-- {
		next := &SkipNode{
			value: math.MinInt32,
			level: i,
		}
		cur.down = next
		cur = next
	}
	return s
}

func (s *SkipList) Search(target int) bool {
	cur := s.topLeft
	for cur != nil {
		if cur.value == target {
			return true
		}
		if cur.right != nil && cur.right.value <= target {
			cur = cur.right
		} else {
			cur = cur.down
		}
	}
	return false
}

func (s *SkipList) Insert(target int) {
	level := s.GenerateRandomLevel()
	cur := s.topLeft
	last := &SkipNode{
		value: target,
		level: level,
	}
	for cur != nil {
		if cur.right != nil && cur.right.value <= target {
			cur = cur.right
		} else {
			if cur.level == level {
				level--
				next := &SkipNode{
					value: target,
					level: level,
				}
				last.right = cur.right
				cur.right = last
				last.down = next
				last = next
			}
			cur = cur.down
		}
	}
}

func (s *SkipList) Erase(target int) {
	cur := s.topLeft
	var real, tmp *SkipNode
	for cur != nil {
		if real == nil && cur.right != nil && cur.right.value == target {
			real = cur.right
		}
		if cur.right != nil && (cur.right.value < target || (real != nil && cur.right != real)) {
			cur = cur.right
		} else {
			if real != nil {
				tmp = real.down
				cur.right = real.right
				real = tmp
			}
			cur = cur.down
		}
	}
}

func (s *SkipList) Print(level int) {
	cur := s.topLeft
	for cur != nil {
		if cur.level <= level {
			fmt.Printf("level: %d >> ", cur.level)
			r := cur
			for r != nil {
				fmt.Printf("%d ", r.value)
				r = r.right
			}
			fmt.Println("")
		}
		cur = cur.down
	}

}

func main() {
	s := Constructor()
	list := []int{6, 1, 2, 10, 4, -1, 1, -10, 5, 11, 1}
	for _, n := range list {
		s.Insert(n)
	}
	s.Print(5)
	//for i := -11; i < 13; i++ {
	//	fmt.Println(i, s.Search(i))
	//}

	s.Erase(1)
	s.Erase(-1)
	s.Erase(-10)
	s.Erase(5)
	s.Print(5)
}
