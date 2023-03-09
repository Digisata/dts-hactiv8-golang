package main

import "fmt"

func main() {
	countMap := map[string]int{}
	for _, v := range "selamat malam" {
		fmt.Println(string(v))
		countMap[string(v)] = countMap[string(v)] + 1
	}
	fmt.Println(countMap)
}
