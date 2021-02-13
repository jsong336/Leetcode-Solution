package main

import (
    "strconv"
    "strings"
    "math"
	"fmt"
)

func reverse(x int) int {
	// Disregard negative
	xAbs := x
	if x < 0 {
		xAbs = -x
	}

	// Convert to string
	xStr := strconv.Itoa(xAbs)
	
	// Reverse string
	xStrArr := strings.Split(xStr, "")
	for i := 0; i < len(xStrArr)/2 ; i++ {
		tmp := xStrArr[i]
		xStrArr[i] = xStrArr[len(xStr) - i - 1]
		xStrArr[len(xStrArr) - i - 1] = tmp
	}

	xRevAbsStr := strings.Join(xStrArr, "")
	xRev, _ := strconv.Atoi(xRevAbsStr)
	if xRev > int(math.Pow(2, 31)){
		// don't like this lol
		xRev = 0
	}
	
	if x < 0{
		xRev = -xRev
	}
	return xRev
}

func main() {
	x := -1534236469
	answer := reverse(x)
	fmt.Println(answer)
}