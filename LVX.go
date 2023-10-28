package main

import (
	"fmt"
	"math/rand"
	"time"
)

func search(target int) {
	left, right := 0, 99
	var (
		i int
		a [100]int
	)
	for i = 0; i < 100; i++ {
		a[i] = i + 1
	}
	for left <= right {
		mid := (left + right) / 2
		if a[mid] == target {
			fmt.Printf("此数在a数组中的位置为a[%d]\n", mid)
			break
		} else if a[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	target2 := rand.Intn(100) + 1
	fmt.Println("此次生成的随机数值为", target2)
	search(target2)
}
