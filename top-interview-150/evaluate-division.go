package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	id := map[string]int{}

	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	graph := make([][]float64, len(id))
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]float64, len(id))
	}

	for i, equation := range equations {
		v, w := id[equation[0]], id[equation[1]]
		graph[v][w] = values[i]
		graph[w][v] = 1.0 / values[i]
	}

	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k] > 0.000001 && graph[k][j] > 0.000001 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, sExist := id[q[0]]
		end, eExist := id[q[1]]

		if !sExist || !eExist || graph[start][end] == 0 {
			ans[i] = -1
		} else {
			ans[i] = graph[start][end]
		}
	}

	return ans
}
