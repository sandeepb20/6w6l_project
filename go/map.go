package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var sl []string = strings.Fields(s)
	for i := 0; i < len(sl); i++ {
		word := sl[i]
		elem, ok := m[word]
		if ok == true {
			m[word] = elem + 1
		} else {
			m[word] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
