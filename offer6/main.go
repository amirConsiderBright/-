package main

import (
	"fmt"
	"strings"
)

type USER struct {
	age  int
	name string
}

func lengthOfLongestSubstring(s string) (a int) {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	left, right := 0, 0
	for _, v := range s {
		index := strings.Index(s[left:right], string(v))
		if index != -1 {
			left += 1 + index
		}
		//fmt.Println(i)
		right++
		if a < right-left {
			a = right - left
		}
	}
	return
}
func main() {
	a := lengthOfLongestSubstring("AbcA")
	fmt.Println(a)
	//fmt.Println(&user)
	bytes := []byte{'h', 'e', 'l', 'l', 'o', '0'}
	reverseString(bytes)
}
func reverseString(s []byte) {
	i2 := len(s)
	i3 := i2 - 1
	for i := 0; i < i2/2; i++ {
		s[i], s[i3-i] = s[i3-i], s[i]
	}
	//fmt.Println(string(s))
}
