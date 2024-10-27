package main

func countReports(numSentCh chan int) int {
	count := 0
	for {
		if v, ok := <-numSentCh; ok {
			count += v
		} else {
			break
		}

	}
	return count
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

