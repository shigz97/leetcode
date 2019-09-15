package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := ""
	n2 := []byte(num2)
	for j := 0; j < len(n2); j++ {
		time := int(n2[j]) - 48
		tmp := ""
		for i := 0; i < time; i++ {
			tmp = add(tmp, num1)
		}

		tp := []byte(tmp)
		for k := j + 1; k < len(n2); k++ {
			tp = append(tp, '0')
		}
		res = add(res, string(tp))
	}
	return res
}

func add(num1, num2 string) string {
	n1 := []byte(reverse(num1))
	n2 := []byte(reverse(num2))
	flag := 0
	res := []byte{}
	var i, j, tp int
	//fmt.Println(n1, n2)
	for i, j = 0, 0; i < len(n1) && j < len(n2); i, j = i+1, j+1 {
		tp = int(n1[i]) + int(n2[j]) - 96 + flag
		//fmt.Println(tp)
		res = append(res, byte(tp%10+48))
		flag = tp / 10
	}
	//fmt.Println(res, i, j)
	if flag == 1 {
		if i == len(n1) && j == len(n2) {
			res = append(res, byte(49))
		} else {
			for ; i < len(n1); i++ {
				tp = flag + int(n1[i]) - 48
				res = append(res, byte(tp%10+48))
				flag = tp / 10
			}
			for ; j < len(n2); j++ {
				tp = flag + int(n2[j]) - 48
				res = append(res, byte(tp%10+48))
				flag = tp / 10
			}
			if flag == 1 {
				res = append(res, byte(49))
			}
		}
	} else {
		res = append(res, n1[i:]...)
		res = append(res, n2[j:]...)
	}
	return reverse(string(res))
}

func reverse(num string) string {
	r := []rune(num)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	num1, num2 := "123", "456"
	fmt.Println(multiply(num1, num2))
}
