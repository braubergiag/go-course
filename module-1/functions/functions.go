package main

import (
	"fmt"
	"sort"
)

func sum(a int, b int) int {
	return a + b
}
func sum_3(a, b, c int) int {
	return a + b + c
}
func sum_all(nums ...int) {
	fmt.Print(nums, " -> ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	a := []int{1, 2, 4, 8, 16, 32, 64, 128}
	x := 53
	closest := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	fmt.Println(a[closest])
}
