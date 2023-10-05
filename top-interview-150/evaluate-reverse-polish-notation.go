package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	symbol := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	var arr []string
	for _, token := range tokens {
		if !symbol[token] {
			arr = append(arr, token)
			continue
		}

		a, _ := strconv.Atoi(arr[len(arr)-2])
		b, _ := strconv.Atoi(arr[len(arr)-1])
		arr = arr[:len(arr)-2]

		var c int
		switch token {
		case "+":
			c = a + b
		case "-":
			c = a - b
		case "*":
			c = a * b
		case "/":
			c = a / b
		}

		arr = append(arr, fmt.Sprint(c))
	}

	res, _ := strconv.Atoi(arr[0])
	return res
}
