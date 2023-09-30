package august

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) == 0 || len(S) == 1 {
		return S
	}

	var bs strings.Builder
	curr := S[0]
	fast := 1
	length := 1
	for ; fast < len(S); fast++ {
		if curr == S[fast] {
			length++
			continue
		}

		bs.WriteByte(curr)
		bs.WriteString(strconv.Itoa(length))
		curr = S[fast]
		length = 1
	}
	bs.WriteByte(curr)
	bs.WriteString(strconv.Itoa(length))

	if bs.Len() >= len(S) {
		return S
	}
	return bs.String()
}
