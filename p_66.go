package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 1 && digits[0] == 0 {
		digits[0]++
		return digits
	}
	digits = reverse(digits)
	plus, advan := 0, 1
	for i := range digits {
		plus = digits[i] + advan
		digits[i] = plus % 10
		advan = plus / 10
	}
	if advan > 0 {
		digits = append(digits, 1)
	}
	return reverse(digits)
}

func reverse(in []int) []int {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}

func main() {
	in := []int{1, 2, 3}
	fmt.Println(plusOne(in))
}
