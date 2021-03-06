// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
func comma(s string) string {
	i := strings.Index(s, ".")
	if i == -1 {
		return leftComma(s)
	}
	return leftComma(s[:i]) + "." + rightComma(s[i+1:])
}

func leftComma(s string) string {
	n := len(s)
	if n <= 3 || (n == 4 && s[0] == '-') {
		return s
	}
	return leftComma(s[:n-3]) + "," + s[n-3:]
}

func rightComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:3] + "," + rightComma(s[3:])
}

//!-
