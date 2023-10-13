package main

func canFinish(numCourses int, prerequisites [][]int) bool {

	inDegree := make(map[int]int, 0)
	preHash := make(map[int][]int, 0)

	for _, prerequisite := range prerequisites {
		cur, pre := prerequisite[0], prerequisite[1]
		inDegree[cur]++
		preHash[pre] = append(preHash[pre], cur)
	}

	zeroDegrees := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			zeroDegrees = append(zeroDegrees, i)
		}
	}

	count := 0
	for len(zeroDegrees) > 0 {
		newDegree := make([]int, 0)
		for _, zeroDegree := range zeroDegrees {
			count++
			for _, item := range preHash[zeroDegree] {
				inDegree[item]--
				if inDegree[item] == 0 {
					newDegree = append(newDegree, item)
				}
			}
		}
		zeroDegrees = newDegree
	}

	return numCourses == count
}
