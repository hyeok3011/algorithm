package main

func solution(name []string, yearning []int, photos [][]string) []int {
	nameScoreMap := make(map[string]int)
	for index := range name {
		nameScoreMap[name[index]] = yearning[index]
	}

	photoScores := make([]int, len(photos))
	for index, photo := range photos {
		total := 0
		for _, name := range photo {
			if score, exist := nameScoreMap[name]; exist {
				total += score
			}
		}
		photoScores[index] = total
	}
	return photoScores
}
