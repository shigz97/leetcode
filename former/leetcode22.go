package main

import "fmt"

var (
	ret []string
	br  = map[int]byte{
		1:  '(',
		-1: ')',
	}
)

func generateParenthesis(n int) []string {
	res := []int{}
	generate(n, 0, res, n*2)
	return ret
}

func generate(x int, sum int, res []int, cap int) {
	if sum < 0 {
		return
	}
	if x == 0 {
		for i := len(res); i < cap; i++ {
			res = append(res, -1)
		}
		//fmt.Println(res)
		//dic = append(dic, res)
		//fmt.Println(res, br)
		str := []byte{}
		for _, k := range res {
			str = append(str, br[k])
		}
		ret = append(ret, string(str))
		return
	}
	res = append(res, 1)
	generate(x-1, sum+1, res, cap)
	res[len(res)-1] = -1
	//fmt.Println(res, x, sum)
	generate(x, sum-1, res, cap)
	return
}

func main() {
	fmt.Println(generateParenthesis(1))
}
