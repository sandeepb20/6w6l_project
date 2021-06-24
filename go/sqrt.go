package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " hi "

	var sl []string = strings.Fields(s)

	m := make(map[string]int)

	for i := 0; i < len(sl); i++ {
		word := sl[i]
		elem, ok := m[word]
		if ok == true {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}

	fmt.Println(m)
}