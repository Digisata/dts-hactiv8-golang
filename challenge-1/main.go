package main

import (
	"fmt"
)

func main() {
	i, j, k := 21, true, 123.456

	fmt.Printf("value = %d\n", i)
	fmt.Printf("type = %T\n", i)
	fmt.Printf("char = %s\n", "%")
	fmt.Printf("bool = %t\n", j)
	fmt.Printf("russia char = %c\n", 'Я')
	fmt.Printf("base 2 = %b\n", i)
	fmt.Printf("base 10 = %d\n", i)
	fmt.Printf("base 10 = %o\n", i)
	fmt.Printf("base 16 = %x\n", 15)
	fmt.Printf("base 16 = %X\n", 3859)
	fmt.Printf("unicode = %U\n", 'Я')
	fmt.Printf("float = %f\n", k)
	fmt.Printf("scientific = %E\n", 1.234560E+02)
}