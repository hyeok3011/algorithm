package main

// https://school.programmers.co.kr/learn/courses/30/lessons/178871
func solution(players []string, callings []string) []string {
	rank := make(map[string]int)
	for index, name := range players {
		rank[name] = index
	}

	for _, calling := range callings {
		currentRank := rank[calling]

		playerSwap(players, currentRank, currentRank-1)
		rank[calling] -= 1
		rank[players[currentRank]] += 1
	}

	return players
}

func playerSwap(players []string, first, second int) {
	players[first], players[second] = players[second], players[first]
}
