package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {


	primaryCost := len(primary)
	secondaryCost := primaryCost + len(secondary)
	tertiaryCost := secondaryCost + len(tertiary)


	return [3]string{primary, secondary, tertiary}, [3]int{primaryCost, secondaryCost, tertiaryCost}

	// ?
}

