package main

import (
	"fmt"
	"strings"
)

func main() {
	const baseString = `Don't communicate by sharing memory, share memory by communicating.
	Concurrency is not parallelism.
	Channels orchestrate; mutexes serialize.
	The bigger the interface, the weaker the abstraction.
	Make the zero value useful.
	interface{} says nothing.
	Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
	A little copying is better than a little dependency.
	Syscall must always be guarded with build tags.
	Cgo must always be guarded with build tags.
	Cgo is not Go.
	With the unsafe package there are no guarantees.
	Clear is better than clever.
	Reflection is never clear.
	Errors are values.
	Don't just check errors, handle them gracefully.
	Design the architecture, name the components, document the details.
	Documentation is for users.
	Don't panic.`

	newStrings := strings.Split(baseString, "\n")

	for _, str := range newStrings {
		s := strings.TrimSpace(str)
		fmt.Printf("\nProcessing %s\n", s)
		countTokens(s)
	}

}

func countTokens(str string) {
	tokenMap := map[rune]int{}

	for _, char := range str {
		tokenMap[char] = tokenMap[char] + 1
	}

	// Print out the values
	for key, value := range tokenMap {
		fmt.Printf("'%c'='%d', ", key, value)
	}
}
