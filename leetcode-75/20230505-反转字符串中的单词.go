package leetcode_75

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slow, fast := 0, 0
	bs := []byte(s)
	// 删除头部冗余空格
	for len(bs) > 0 && fast < len(bs) && bs[fast] == ' ' {
		fast++
	}
	// 删除单词间冗余空格
	for ; fast < len(bs); fast++ {
		if fast-1 > 0 && bs[fast-1] == bs[fast] && bs[fast] == ' ' {
			continue
		}
		bs[slow] = bs[fast]
		slow++
	}
	// 删除尾部冗余空格
	if slow-1 > 0 && bs[slow-1] == ' ' {
		bs = bs[:slow-1]
	} else {
		bs = bs[:slow]
	}
	// 2.反转整个字符串
	reverse(&bs, 0, len(bs)-1)
	// 3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for ; i < len(bs); i++ {
		j := i
		for ; j < len(bs) && bs[j] != ' '; j++ {
		}
		reverse(&bs, i, j-1)
		i = j
	}
	return string(bs)
}

func reverse(a *[]byte, i, j int) {
	for i < j {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		i++
		j--
	}
}
