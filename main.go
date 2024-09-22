package main

func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body) * 2
	}
	return len(e.body) * 5
}

func (e email) format() string {
	var isSubscribedText string
	if e.isSubscribed {
		isSubscribedText = "Subscribed"
	} else {
		isSubscribedText = "Not Subscribed"
	}

	return "'"+e.body+"' | "+isSubscribedText
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

