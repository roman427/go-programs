package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		if len(s)-1 < 4 {
			return s
		}
		buf.WriteByte(s[0])

		for i, v := range s[1:] {
			buf.WriteRune(v)
			if (i == 0) || (i%3 == 0) {
				buf.WriteByte(',')
			}
		}
	} else {
		if len(s) < 4 {
			return s
		}

		for i, v := range s {
			buf.WriteRune(v)
			if (i == 0) || (i%3 == 0) {
				buf.WriteByte(',')
			}
		}
	}

	if strings.HasSuffix(buf.String(), ",") {
		temp := buf.String()
		temp = temp[:len(temp)-1]
		return temp
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("12"))
}
