package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func bisect(num int) (x, y int) {
	x = num / 2
	y = x
	return
}

func create(prefix string) func(string) {
	return func(msg string) {
		fmt.Println(prefix + msg)
	}
}

func main() {
	fmt.Println(add(1, 2))
	x, y := swap(5, 10)
	fmt.Println(x, y)
	a, b := bisect(19)
	fmt.Println(a, b)

	cat := create("🐈 ")
	cat("I'm a cat")
}
