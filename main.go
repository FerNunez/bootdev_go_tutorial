package main

func bulkSend(numMessages int) float64 {
	total :=0.0

	for i:=1; i<=numMessages; i++{
		total += 1.0 + (float64(i-1)* 0.01)
	}

	return total
}

