package main

import "time"

func processMessages(messages []string) []string {
	// ?

	output := make([]string, 0)
	ch := make(chan string)
	for _, m := range messages {
		go func(ch chan string, m string) {

			ch <- process(m)
		}(ch, m)
	}

	count := 0
	for i := range ch {
		println("waiting to read!")
		output = append(output, i)

		count += 1
		if count >= len(messages) {
			break
		}

	}
	close(ch)
	return output
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

