package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?


	for idxMsg, word := range msg{
		for _, badWord := range badWords{
			if word == badWord{
				return idxMsg
			}
		}
	}

	return -1
}

