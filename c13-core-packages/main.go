package main

import (
	"fmt"
	"strings"
)

func main() {
	stringsDemo()
}

func stringsDemo() {
	testStr := "This is my test string"
	fmt.Println(
		strings.Contains(testStr, "test"),
		strings.Count(testStr, "t"),
		strings.HasPrefix(testStr, "Th"),
		strings.HasSuffix(testStr, "ng"),
		strings.Index(testStr, "i"),
		strings.Join(strings.Split(testStr, " "), "-"),
		strings.Repeat("a", 10),
		strings.Replace(testStr, "This is", "That was", 1),
		strings.ToLower(testStr),
		strings.ToUpper(testStr),
	)
}