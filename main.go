package main

func addEmailsToQueue(emails []string) chan string {

	channel := make(chan string, len(emails))

	for _,e := range emails{
		channel <- e
	}

	return channel
}

