package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {

	words := make(map[string]int)
	total := 0

	for _, message := range messages {

		for _, word := range strings.Fields(message) {

			lowerWorld := strings.ToLower(word)
			if _, ok := words[lowerWorld]; !ok {

				words[lowerWorld] = 0
				total += 1
			}

		}
	}
	return total
}

