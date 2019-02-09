package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

// WordCount function
func WordCount(s string) map[string]int {
	m := map[string]int{}

	for _, word := range strings.Fields(s) {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
