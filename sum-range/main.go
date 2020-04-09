package main

import (
	"fmt"
	"os"
	"strconv"
)

//sumRange returns a sum of integers from start till end, checking all errors
func sumRange(start, end string) int {
	var sum int

	s, err := strconv.Atoi(start)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	e, err := strconv.Atoi(end)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}

	for i := s; i < e; i++ {
		sum += i
	}

	return sum
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: ./%s start end\n", os.Args[0])
		os.Exit(1)
	}

	sum := sumRange(os.Args[1], os.Args[2])

	fmt.Fprintf(os.Stdout, "Sum of numbers from %s till %s is: %d\n", os.Args[1], os.Args[2], sum)
}
