package main

func canCompleteCircuit(gas []int, cost []int) int {
	currentGas := 0
	totalGas := 0
	index := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			currentGas = 0
			index = i + 1
		}
	}

	if totalGas < 0 {
		return -1
	}
	return index
}
