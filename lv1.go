package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	var x, y int
	fmt.Scanf("%d,%d", &x, &y)
	result := add(x, y)
	fmt.Println(result)
}
