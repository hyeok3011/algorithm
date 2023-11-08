package main

import (
	"sort"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/42587
// 간단하게 queue로 풀었다.

type Task struct {
	Priority int
	Index    int
}

func solution(priorities []int, location int) int {
	queue := Queue{}
	for index, v := range priorities {
		queue.Push(&Task{
			Index:    index,
			Priority: v,
		})
	}
	sort.Slice(priorities, func(i, j int) bool {
		return priorities[i] > priorities[j]
	})

	executionOrder := 0
	priorityIndex := 0
	for {
		task := queue.Pop()
		if priorities[priorityIndex] == task.Priority {
			executionOrder++
			priorityIndex++
			if task.Index == location {
				break
			}
		} else {
			queue.Push(task)
		}
	}
	return executionOrder
}

type Queue struct {
	v []*Task
}

func (q *Queue) Push(task *Task) {
	q.v = append(q.v, task)
}

func (q *Queue) Pop() *Task {
	task := q.v[0]
	q.v = q.v[1:]
	return task
}

func (q *Queue) IsEmpty() bool {
	return len(q.v) == 0
}

// 다른 사람의 풀이 방법이다.
// https://school.programmers.co.kr/learn/courses/30/lessons/42587/solution_groups?language=go
func otherSolution(p []int, loc int) int {
	lenP := len(p)
	pDesc := make([]int, lenP)
	copy(pDesc, p)

	sort.Sort(sort.Reverse(sort.IntSlice(pDesc)))

	printed := 0
	max := lenP * (lenP + 1) / 2

	for i := 0; i < max; i++ {
		curLoc := i % lenP
		cur := p[curLoc]
		if cur == -1 || cur != pDesc[0] {
			continue
		}

		printed++
		p[curLoc] = -1
		pDesc = pDesc[1:]

		if curLoc == loc {
			break
		}
	}

	return printed
}
