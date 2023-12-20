package codewar

import "regexp"

// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/solutions/go/all/clever

func Solution(str string) []string {
	if len(str)%2 == 1 {
		str = str + "_"
	}
	r := make([]string, 0, len(str)/2)
	for i := 0; i < len(str)/2; i++ {
		r = append(r, str[i*2:(i+1)*2])
	}
	return r
}

func Solution2(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}
