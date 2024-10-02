package main

import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
	for i:=0; i<len(messages); i++{

		tag := tagger(messages[i])

		messages[i].tags = tag
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}

	if strings.Contains(msg.content, "urgent") || strings.Contains(msg.content, "Urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(msg.content, "sale") || strings.Contains(msg.content, "Sale") {
		tags = append(tags, "Promo")
	}
	return tags
	// ?
}

