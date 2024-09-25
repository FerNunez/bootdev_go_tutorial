package main

func maxMessages(thresh int) int {


	maxNumberMessages:=0
	total := 0
	for i:=0; total<=thresh; i++{
		total+= 100+i
		maxNumberMessages = i
	}
	return maxNumberMessages
	
}

