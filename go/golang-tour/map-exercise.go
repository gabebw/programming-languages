// http://tour.golang.org/#41
package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	fields := strings.Fields(s)
	for i := 0; i < len(fields); i++ {
		word := fields[i]
		result[word] = result[word] + 1
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
