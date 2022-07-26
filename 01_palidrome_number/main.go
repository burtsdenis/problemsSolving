package main

import (
	"fmt"
)

func main() {
	fmt.Println(palindrome(123)) // false
	fmt.Println(palindrome(121)) // true
	fmt.Println(palindrome(1412213)) // false
	fmt.Println(palindrome(125521)) // true
}

func palindrome(x int) bool {
	 revNum := 0
	 tempNum := x
	 for tempNum > 0 {
		 r := tempNum % 10
		 revNum = revNum * 10 + r
		 tempNum /= 10
	 }

	return x == revNum
}