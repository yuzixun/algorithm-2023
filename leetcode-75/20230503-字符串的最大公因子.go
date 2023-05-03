package leetcode_75

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	l1, l2 := len(str1), len(str2)
	return str1[:gcd(l1, l2)]
}

func gcd(i, j int) int {
	if j == 0 {
		return i
	}
	return gcd(j, i%j)
}
