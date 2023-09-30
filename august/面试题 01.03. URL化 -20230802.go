package august

func replaceSpaces(S string, length int) string {
	counter := len(S) - 1
	bs := []byte(S)
	for i := length - 1; i >= 0; i-- {
		if S[i] == ' ' {
			bs[counter] = '0'
			counter--
			bs[counter] = '2'
			counter--
			bs[counter] = '%'
		} else {
			bs[counter] = S[i]
		}
		counter--
	}
	return string(bs[counter+1:])
}
