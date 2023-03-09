package main

import "fmt"

func main() {
	i, j, counter := 0, 0, 0
	for i < 5 || j <= 10 {
		if i <= 4 {
			fmt.Printf("Nilai i = %d\n", i)
			i++
			continue
		}

		if j == 5 {
			for _, v := range "САШАРВО" {
				fmt.Printf("character %U '%s' starts at byte position %d\n", v, string(v), counter)
				counter += 2
			}
			j++
			continue
		}

		fmt.Printf("Nilai j = %d\n", j)
		j++
	}
}
