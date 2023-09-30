package august

func isUnique(astr string) bool {
	a := 0
	for _, i := range astr {
		if (a & (1 << (i - 'a'))) != 0 {
			return false
		}

		a |= 1 << (i - 'a')
	}
	return true
}
