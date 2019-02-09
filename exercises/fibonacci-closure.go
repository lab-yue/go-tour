package main

import "fmt"

func fibonacci() func() []int {
	arr := []int{0, 1}
	return func() []int {
		arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
		return arr
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
