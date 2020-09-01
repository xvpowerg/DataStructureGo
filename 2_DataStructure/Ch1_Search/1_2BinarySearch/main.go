package main

import (
	"fmt"
	"sort"
)

func main() {
	array := [10]int{5, 8, 13, 15, 2, 7, 3, 11, 4, 18}
	isFind := false
	keyWord := 15
	sort.Ints(array)
	if !isFind {
		fmt.Printf("Not Find!")
	}
}
