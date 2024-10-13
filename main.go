package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {

	newSuggestions := make(map[string]bool)
	alreadyFriend := make(map[string]bool)

	// get username's friends
	if userFriends, ok := friendships[username]; ok {

		// add already friend into a map/hashset
		for _, userFriendName := range userFriends {
			alreadyFriend[userFriendName] = false
		}

		// loop friends friends
		for _, userFriendName := range userFriends {

			if friendsOfFriends, okThereIsMap := friendships[userFriendName]; okThereIsMap {

				for _, friendsOfFriendName := range friendsOfFriends {

					isNotUser := friendsOfFriendName != username
					_, okAlreadyFriend := alreadyFriend[friendsOfFriendName]
					_, okNewSuggestions := newSuggestions[friendsOfFriendName]

					if isNotUser && !okAlreadyFriend && !okNewSuggestions {
						newSuggestions[friendsOfFriendName] = true
					}

				}

			}

		}

	}

	if numbSugestions := len(newSuggestions); numbSugestions > 0 {

		v := make([]string, 0, len(newSuggestions))

		for key := range newSuggestions {
			v = append(v, key)
		}

		return v

	}

	return nil

}
