package main

import "fmt"

func Calculator(a, b int) int {
	return a + b
}

func main() {
	result := Calculator(3, 4)
	fmt.Println(result) // 输出：7
}
