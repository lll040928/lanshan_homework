package main

import "fmt"

func prime(x int) int {
	var i int
	for i = 2; i < x; i++ {
		if x%i == 0 {
			return -1
			break
		}
	}
	return 0
}
func main() {
	var x int
	fmt.Scanln(&x)
	isprime := prime(x)
	if isprime == 0 {
		fmt.Println("是素数")
	} else {
		fmt.Println("不是素数")
	}
}
