package main

func countConnections(groupSize int) int {

	total := 0
	for i := 1; i <= groupSize; i++ {

		for j := i+1; j <= groupSize; j++ {
			total += 1
		}
	}
	return total
}
