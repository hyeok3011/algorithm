package main

func solution(bridge_length int, weight int, truck_weights []int) int {
	truckWeightIndex := 0
	duration := 0
	currentBridgheWeight := 0
	truckQueue := []int{}
	timeQueue := []int{}
	for truckWeightIndex < len(truck_weights) {
		duration += 1
		if len(timeQueue) > 0 && timeQueue[0] == duration {
			currentBridgheWeight -= truckQueue[0]
			timeQueue = timeQueue[1:]
			truckQueue = truckQueue[1:]
		}

		if currentBridgheWeight+truck_weights[truckWeightIndex] <= weight && len(truckQueue) < bridge_length {
			timeQueue = append(timeQueue, duration+bridge_length)
			truckQueue = append(truckQueue, truck_weights[truckWeightIndex])
			currentBridgheWeight += truck_weights[truckWeightIndex]
			truckWeightIndex += 1
		} else {
			duration = timeQueue[0] - 1
		}

	}

	duration += bridge_length

	return duration
}
