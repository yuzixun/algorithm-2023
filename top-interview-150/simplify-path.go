package main

import "strings"

func simplifyPath(path string) string {

	stacks := strings.Split(path, "/")

	var arr []string
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}

		if stack == "." {
			continue
		}

		if stack == ".." {
			if len(arr) > 0 {
				arr = arr[:len(arr)-1]
			}

			continue
		}
		arr = append(arr, stack)
	}

	return "/" + strings.Join(arr, "/")
}
