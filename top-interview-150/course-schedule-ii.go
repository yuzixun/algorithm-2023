package main

func findOrder(numCourses int, prerequisites [][]int) []int {

	inDegree := make(map[int]int, 0)
	preHash := make(map[int][]int, 0)
	for _, prerequisite := range prerequisites {
		inDegree[prerequisite[0]]++
		preHash[prerequisite[1]] = append(preHash[prerequisite[1]], prerequisite[0])
	}

	zeroDegrees := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			zeroDegrees = append(zeroDegrees, i)
		}
	}

	result := make([]int, 0, numCourses)
	for len(zeroDegrees) > 0 {

		newDegrees := make([]int, 0)
		for _, item := range zeroDegrees {
			result = append(result, item)

			for _, cur := range preHash[item] {
				inDegree[cur]--
				if inDegree[cur] == 0 {
					newDegrees = append(newDegrees, cur)
				}
			}
		}

		zeroDegrees = newDegrees
	}

	if len(result) == numCourses {
		return result
	}
	return []int{}
}
