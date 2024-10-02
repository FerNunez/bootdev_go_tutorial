package main

func getCounts(messagedUsers []string, validUsers map[string]int) {

for _, mu := range messagedUsers{
	if _, ok := validUsers[mu]; ok {

			validUsers[mu]+=1

		}
	}
}

