package main

import "fmt"

func balancedString(s string) int {
	ba := len(s) / 4
	var count [4]int
	for _, k := range []byte(s) {
		switch k {
		case 'Q':
			count[0]++
		case 'W':
			count[1]++
		case 'E':
			count[2]++
		case 'R':
			count[3]++
		}
	}
	res := 0

}

func main() {
	s := "QQQQ"
	fmt.Println(balancedString(s))
}
