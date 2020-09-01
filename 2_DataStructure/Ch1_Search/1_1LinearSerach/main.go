package main

import "fmt"

func main() {
	array := [10]int{5, 8, 13, 15, 2, 7, 3, 11, 4, 18}
	isFind := false
	keyWord := 15
	for i, v := range array {
		if v == keyWord {
			fmt.Printf("index:%d value:%d", i, v)
			isFind = true
			break
		}
	}
	if !isFind {
		fmt.Printf("Not Find!")
	}
}
