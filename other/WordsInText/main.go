package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Hello, wolrd!"

	regex := regexp.MustCompile("\\w+")

	words := regex.FindAllString(text, -1)

	fmt.Println(words)

}
