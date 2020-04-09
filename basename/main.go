//program discards / and .extension from a filename
package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	s := basenamePath(os.Args[1])
	fmt.Printf("path: %s\n", s)

	s = basenameStrings(os.Args[1])
	fmt.Printf("strings: %s\n", s)
}

func basenameStrings(s string) string {
	slash := strings.LastIndex(s, "/")
	if slash != -1 {
		s = s[slash+1:]
	}
	if dot := strings.LastIndex(s, "."); dot != -1 {
		s = s[:dot]
	}

	return s
}

func basenamePath(s string) string {
	return path.Base(s)
}
