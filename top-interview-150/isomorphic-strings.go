package main

func isIsomorphic(s string, t string) bool {

	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sData, tData := s[i], t[i]

		if (s2t[sData] > 0 && s2t[sData] != tData) || (t2s[tData] > 0 && t2s[tData] != sData) {
			return false
		}
		s2t[sData] = tData
		t2s[tData] = sData
	}
	return true
}
