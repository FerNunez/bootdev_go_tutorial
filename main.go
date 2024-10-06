package main

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here

	result := make(map[rune]map[string]int)
	for _, name := range names {

		runes := []rune(name)

		

		if _, ok := result[runes[0]];!ok {
			result[runes[0]] = make(map[string]int)

		}

		result[runes[0]][name] += 1
	}
	return result
}

