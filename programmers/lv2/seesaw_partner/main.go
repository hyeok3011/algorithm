package main

// https://school.programmers.co.kr/learn/courses/30/lessons/152996
// 2:2, 2:3, 2:4, 3:2, 3:3, 3:4, 4:2, 4:3, 4:4
// 위 비율을 고민해야한다. 위 비율을 다시 정리하면
// 1:1, 2:3, 1:2, 3:4만 고민하면 된다.
// 길이가 10만인것을 보고 2중 for문은 죽었다가 깨도 안되겠구만 생각하고
// 무개별 인원수 map을 만들고 개수만큼만 더하는 식으로 처리했더니 문제 없이 잘 풀렸다.
// 처음에 weightCounter를 map[int]int로 만들었더니 계속 에러나서 못찾고 있었다.
// 이런 낚시 조심하자...
func solution(weights []int) int64 {
	var partners int64
	weightCounter := make(map[float32]int)

	for i := range weights {
		if _, exist := weightCounter[float32(weights[i])]; exist {
			weightCounter[float32(weights[i])]++
		} else {
			weightCounter[float32(weights[i])] = 1
		}
	}

	for weight, count := range weightCounter {
		// 1:1
		if count > 1 {
			partners += int64(count * (count - 1) / 2)
		}

		// 2:3
		if v, exist := weightCounter[weight*3/2]; exist {
			partners += int64(count * v)
		}

		// 1:2
		if v, exist := weightCounter[weight*2]; exist {
			partners += int64(count * v)
		}

		// 3:4
		if v, exist := weightCounter[weight*4/3]; exist {
			partners += int64(count * v)
		}

	}
	return partners
}
