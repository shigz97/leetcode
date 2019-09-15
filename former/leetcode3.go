package main

import "fmt"

func main() {
	str := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var max, i int
	i = 0
	//max = 1
	dic := make(map[byte]int, len(s))
	for j, ch := range []byte(s) {
		if index, ok := dic[ch]; ok && index >= i {
			//fmt.Printf("j=%d,ch=%d\n", index, ch)
			i = index + 1
		}
		dic[ch] = j
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}
