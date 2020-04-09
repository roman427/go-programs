package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func printRunesRange(min, max int) {
	for i := 0; i+min < max; i++ {
		fmt.Printf("%d: %q\n", i, rune(i+min))
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: ./program min max")
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	if max-min <= 0 {
		log.Fatalf("max < min.\n")
	} else {
		printRunesRange(min, max)
	}
}
