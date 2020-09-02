package main

import "fmt"

func naiveStringSearch(strValue string, key string) int {
	count := 0
	keyLen := len(key)
	for i := 0; i < len(strValue); i++ {
		for k := 0; k < len(key); k++ {
			if key[k] != strValue[i+k] {
				break
			}
			if k == keyLen-1 {
				count++
			}
		}
	}
	return count
}
func naiveStringSearchMap(strValue string, key string) int {
	myMap := make(map[string]int)
	myMap[key] = 0
	keyLen := len(key)
	strLen := len(strValue)
	for i := range strValue {
		if i+keyLen <= strLen {
			//fmt.Println(strValue[i : i+keyLen])
			cehckKey := strValue[i : i+keyLen]
			_, exis := myMap[cehckKey]
			if exis {
				myMap[cehckKey]++
			}
		}
	}

	return myMap[key]
}
func main() {
	//count := naiveStringSearch("kewi aapply cherry apple", "ap")
	//fmt.Println(count)
	count2 := naiveStringSearchMap("kewi aapply cherry apple", "ap")
	fmt.Println(count2)
}
