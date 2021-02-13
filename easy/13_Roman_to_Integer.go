package main

import "fmt"

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
    romanDict := map[string]int{
        "I": 1,
        "V":  5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    n_total := 0 
    n_prev := romanDict[string(s[0])]
    
    var n_cur int    
    for _, c := range s[1:] {
        n_cur, _ = romanDict[string(c)]
		
        if n_prev < n_cur {
            n_total = n_total - n_prev
        } else{
            n_total = n_total + n_prev
        }
        n_prev = n_cur
    }
    n_total = n_total + n_prev
    return n_total
}

func main() {
	answer := romanToInt("MCMXCIV")	
	fmt.Printf("answer: %d\n", answer)
}