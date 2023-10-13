package main

var (
	m = map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
)

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = make([]string, 0)
	letterCombinationsHandler(digits, 0, "")

	return res
}

func letterCombinationsHandler(digits string, index int, str string) {
	if len(digits) == index {
		res = append(res, str)
	} else {
		letters := m[int(digits[index]-'0')]

		for _, letter := range letters {
			letterCombinationsHandler(digits, index+1, str+letter)
		}
	}
}
