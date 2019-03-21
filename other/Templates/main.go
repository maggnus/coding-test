package main

import (
	"fmt"
	"regexp"
	"strings"
)

func fillTemplate(s string, m map[string]string) string {
	for key := range m {
		r := regexp.MustCompile("\\$?<" + key + "\\$?>")
		b := r.ReplaceAllFunc([]byte(s), func(k []byte) []byte {
			t := string(k)
			if strings.HasPrefix(t, "$<") || strings.HasPrefix(t, "$>") {
				return k
			}
			key = strings.TrimPrefix(t, "<")
			key = strings.TrimSuffix(key, ">")
			return []byte(m[key])
		})
		s = string(b)
	}
	return s
}

func main() {

	s := "As computers <word1> more and more $<powerful$>, they won’t be substitutes for <word2>: they’ll be complements."
	m := map[string]string{
		"word1": "become",
		"word2": "humans",
	}

	result := fillTemplate(s, m)

	fmt.Printf("Input: %s\n", s)
	fmt.Printf("Result: %s\n", result)
}
