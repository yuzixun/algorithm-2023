package main

func addBinary(a string, b string) string {
	aLen, bLen := len(a)-1, len(b)-1
	rLen := Max(aLen, bLen)
	result := make([]byte, rLen+1)

	carry := 0
	for aLen >= 0 && bLen >= 0 {
		carry += int(a[aLen]-'0') + int(b[bLen]-'0')
		result[rLen] = '0' + uint8(carry%2)
		carry /= 2

		aLen--
		bLen--
		rLen--
	}

	for aLen >= 0 {
		carry += int(a[aLen] - '0')
		result[rLen] = '0' + uint8(carry%2)
		carry /= 2

		aLen--
		rLen--
	}

	for bLen >= 0 {
		carry += int(a[aLen] - '0')
		result[rLen] = '0' + uint8(carry%2)
		carry /= 2

		bLen--
		rLen--
	}

	if carry == 1 {
		result = append([]byte{'1'}, result...)
	}
	if len(result) > 1 && result[0] == '0' {
		result = result[1:]
	}

	return string(result)
}
