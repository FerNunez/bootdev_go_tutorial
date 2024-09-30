package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {

	result := make([]float64, 1)

	for i := 0; i < len(costs); i++ {
		day := costs[i].day

		for {
			if day > len(result)-1 {
				result = append(result, 0.)
			} else {
				break
			}
		}

		result[day] += costs[i].value
	}

	return result
}

