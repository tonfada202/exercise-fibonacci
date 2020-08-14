package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	strParts := strings.Split(s, " ")
	m := make(map[string]int)
	for _, key := range strParts {
		d, ok := m[key]
		if ok {
			m[key] = d + 1
		} else {
			m[key] = 1
		}
	}
	return m
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println("---- WordCount ----")
	for index, value := range w {
		fmt.Println(index, "=", value)
	}
	fmt.Printf("***** total = %d ***** \n ", len(w))
}
