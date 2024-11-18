package main

import "fmt"

func main() {
	const name, age = "Kim", 22
	// two ways of doing the same thing
	// %s is for string, %d is for number
	// %v is for value and %T is for type
	fmt.Printf("name: %s age: %d.\n", name, age)
	fmt.Printf("name: %v age: %v.\n", name, age)
	fmt.Printf("name type: %T age type: %T.\n", name, age)
}
