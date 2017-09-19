package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	strlists := strings.Fields(s)
	
	for _,v := range strlists {
		_, flg := m[v]
		if flg {
		  m[v] = m[v] + 1	
		} else {
		  m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

