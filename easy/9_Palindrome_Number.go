package main

import (
	"fmt"
	"strconv"
)
func isPalindromeStr(xStr string) bool {
	if len(xStr) > 1{
		if xStr[0] == xStr[len(xStr) -1] {
			return isPalindromeStr(xStr[1:len(xStr)-1])
		} else {
			return false
		}		
	}
	return true
}
func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	return isPalindromeStr(xStr)
}

func main() {
	x := 1000001
	answer := isPalindrome(x)
	fmt.Println(answer)
}