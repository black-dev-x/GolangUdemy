package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("a:%v \t %b \t %#x \n", a, a, a)
	fmt.Printf("b:%v \t %b \t %#x \n", b, b, b)
	fmt.Printf("c:%v \t %b \t %#x \n", c, c, c)
	fmt.Printf("d:%v \t %b \t %#x \n", d, d, d)
	fmt.Printf("e:%v \t %b \t %#x \n", e, e, e)
	fmt.Printf("f:%v \t %b \t %#x \n", f, f, f)
}
