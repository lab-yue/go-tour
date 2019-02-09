package main

import "fmt"

type dot struct {
	x, y int
}

func main() {
	usePointer()
	useStruct()
	useArray()
	useMap()
}

func usePointer() {
	i := 100
	p := &i
	fmt.Println(*p)
	*p = 2000
	fmt.Println(i)
}

func useStruct() {
	newDot := dot{1, 2}
	fmt.Println(newDot.x)
	p := &newDot
	fmt.Println(p.x, (*p).y)
}

func useArray() {
	var a [2]string
	a[0] = "new"
	a[1] = "array"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[:4]
	fmt.Println(s)
	s = s[:5]
	fmt.Println(s)

	newArray := make([]int, 10)
	newArray = append(newArray, 1, 2, 3, 4)
	fmt.Println(newArray)
	for i, v := range newArray {
		fmt.Println(i, "=>", v)
	}
}

type color struct {
	a, b, c int
}

func useMap() {
	m := map[string]color{
		"#000": {0, 0, 0},
		"#fff": {255, 255, 255},
	}
	delete(m, "#fff")

	if white, ok := m["#fff"]; ok {
		fmt.Println(white)
	} else {
		fmt.Println("Not exist!")
	}
	fmt.Println(m["#000"])
}
