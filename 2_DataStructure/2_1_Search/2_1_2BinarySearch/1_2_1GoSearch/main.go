package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{5, 8, 13, 15, 2, 7, 3, 11, 4, 18}
	sort.Ints(array)
	key := 8
	index := sort.SearchInts(array, key)
	if index < len(array) && array[index] == key {
		fmt.Println("find !")
	} else {
		fmt.Println("not find !")
	}

}
