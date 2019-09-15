package main

import "fmt"

func main() {
	num1 := []int{1, 2}
	num2 := []int{3, 4}
	num := make([]int, 4)
	copy(num, num1)
	//num = append(num, num2...)
	copy(num[len(num2):], num2)
	fmt.Printf("vec=%v", num)
}
