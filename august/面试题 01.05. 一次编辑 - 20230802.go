package august

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)

	if diff := m - n; diff > 1 || diff < -1 {
		return false
	}

	i, j, used := 0, 0, false
	for i < m && j < n {
		if first[i] == second[j] {
			i++
			j++
			continue
		}
		if used {
			return false
		}
		used = true
		switch {
		case m == n:
			i++
			j++
		case m < n:
			j++
		case m > n:
			i++
		}

	}
	return true
}
