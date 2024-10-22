package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?
// When your function receives a pointer to a struct, you might try to access a field like this and encounter an error:
//		msgTotal := *analytics.MessagesTotal
// Instead, use this simple approach using a selector expression.
//		msgTotal := analytics.MessagesTotal
// This approach is shorthand for (*analytics).MessagesTotal and is the recommended, simplest way to access struct fields in Go.


func getMessageText(a *Analytics, m Message) {
	if m.Success {
		a.MessagesSucceeded += 1
	} else {
		a.MessagesFailed += 1
	}

	a.MessagesTotal += 1
}

