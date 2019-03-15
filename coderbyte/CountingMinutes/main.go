package main

import (
	"fmt"
	"strings"
	"time"
)

func CountingMinutes(str string) string {

	// Create midnight time as point of counting
	midnight, _ := time.Parse(time.Kitchen, "12:00AM")

	arr := strings.Split(strings.ToUpper(str), "-")

	start, err := time.Parse(time.Kitchen, arr[0])
	if err != nil {
		panic(err)
	}

	end, err := time.Parse(time.Kitchen, arr[1])
	if err != nil {
		panic(err)
	}

	t1 := start.Sub(midnight).Minutes()
	t2 := end.Sub(midnight).Minutes()

	dur := t2 - t1
	if dur < 0 {
		dur = dur + 24*60 // minutes in 24 hours
	}

	str = fmt.Sprint(dur)

	return str

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(CountingMinutes(readline()))

}
